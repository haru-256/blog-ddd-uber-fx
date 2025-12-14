package presentation_test

import (
	"io"
	"log/slog"
	"testing"

	"github.com/haru-256/blog-ddd-uber-fx/internal/presentation"
	"github.com/haru-256/blog-ddd-uber-fx/internal/presentation/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestModule(t *testing.T) {
	t.Run("Provide Server", func(t *testing.T) {
		app := fxtest.New(t,
			presentation.Module,
			fx.NopLogger,
			// Suppress logger output from infrastructure/server
			fx.Decorate(func() (*slog.Logger, error) {
				return slog.New(slog.NewTextHandler(io.Discard, nil)), nil
			}),
			fx.Supply(
				fx.Annotated{Name: "serverPort", Target: "0"},
				fx.Annotated{Name: "logLevel", Target: "INFO"},
				fx.Annotated{Name: "logFormat", Target: "json"},
				fx.Annotated{Name: "addSource", Target: true},
			),
			// Invokeで検証
			fx.Invoke(fx.Annotate(
				func(
					s *server.ServiceServer,
					routes []server.RouteRegistrar,
				) {
					assert.NotNil(t, s)
					assert.Len(t, routes, 2)
				},
				fx.ParamTags("", `group:"routes"`),
			)),
		)
		app.RequireStart().RequireStop()
	})
}
