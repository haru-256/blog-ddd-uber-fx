package server

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

const (
	healthPath = "/health"
)

type ServiceServerConfig struct {
	Port string
}

func NewServiceServerConfig(serverPort string) *ServiceServerConfig {
	return &ServiceServerConfig{
		Port: serverPort,
	}
}

type ServiceServer struct {
	logger *slog.Logger // ロガー
	e      *echo.Echo   // Echoインスタンス
}

func NewServiceServer(
	cfg *ServiceServerConfig,
	logger *slog.Logger,
	userHandler *UserServiceHandler,
	taskHandler *TaskServiceHandler,
) *ServiceServer {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogError:   true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			latencyMsec := v.Latency.Seconds() * 1000
			if v.Error == nil {
				logger.LogAttrs(c.Request().Context(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("method", v.Method),
					slog.Float64("latency_msec", latencyMsec),
				)
			} else {
				logger.LogAttrs(c.Request().Context(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("method", v.Method),
					slog.Float64("latency_msec", latencyMsec),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == healthPath
		},
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			if len(reqBody) == 0 {
				return
			}
			logger.InfoContext(c.Request().Context(), "req_body", slog.String("body", string(reqBody)))
			if len(resBody) == 0 {
				return
			}
			logger.InfoContext(c.Request().Context(), "res_body", slog.String("body", string(resBody)))
		},
	}))

	// ルーティングの設定
	// 1. ヘルスチェック用エンドポイント
	e.GET(healthPath, func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "healthy",
		})
	})

	// 2. User Service Endpoints
	users := e.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.GET("", userHandler.GetAllUsers)
	users.GET("/:id", userHandler.GetUserById)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)

	// 3. Task Service Endpoints
	tasks := e.Group("/tasks")
	tasks.POST("", taskHandler.CreateTask)
	tasks.GET("", taskHandler.GetAllTasks)
	tasks.GET("/:id", taskHandler.GetTaskById)
	tasks.PUT("/:id", taskHandler.UpdateTask)
	tasks.DELETE("/:id", taskHandler.DeleteTask)

	return &ServiceServer{
		logger: logger,
		e:      e,
	}
}

// RegisterLifecycleHooks はサーバーのライフサイクルフックを登録します。
//
// Parameters:
//   - lc: fxライフサイクル
//   - server: CQRSServiceServer
//   - cfg: サーバー設定
func RegisterLifecycleHooks(lc fx.Lifecycle, server *ServiceServer, cfg *ServiceServerConfig) {
	lc.Append(fx.Hook{
		// AppがStartするときに実行される
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", ":"+cfg.Port)
			if err != nil {
				return err
			}
			server.e.Listener = ln
			go func() {
				server.logger.InfoContext(ctx, "Starting Service Server", slog.String("addr", server.e.Listener.Addr().String()))
				if err = server.e.Start(""); err != nil && !errors.Is(err, http.ErrServerClosed) {
					server.logger.ErrorContext(ctx, "Failed to start server", "error", err)
				}
			}()
			return nil
		},
		// AppがStopするときに実行される
		OnStop: func(ctx context.Context) error {
			server.logger.InfoContext(ctx, "Shutting down Service Server...")
			return server.e.Shutdown(ctx)
		},
	})
}
