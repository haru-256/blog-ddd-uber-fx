package infrastructure_test

import (
	"io"
	"log/slog"
	"testing"

	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/repositories"
	"github.com/haru-256/blog-ddd-uber-fx/internal/infrastructure"
	"github.com/haru-256/blog-ddd-uber-fx/internal/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestModule(t *testing.T) {
	t.Run("Provide Repositories", func(t *testing.T) {
		app := fxtest.New(t,
			infrastructure.Module,
			// テスト時のみデコレートして差し替える
			fx.Decorate(func() (*slog.Logger, error) {
				return slog.New(slog.NewTextHandler(io.Discard, nil)), nil
			}),
			fx.Supply(
				fx.Annotated{Name: "logLevel", Target: "INFO"},
				fx.Annotated{Name: "logFormat", Target: "json"},
				fx.Annotated{Name: "addSource", Target: true},
			),
			fx.Invoke(func(
				userRepo repositories.UserRepository,
				taskRepo repositories.TaskRepository,
			) {
				assert.NotNil(t, userRepo)
				assert.NotNil(t, taskRepo)
				assert.IsType(t, &db.UserRepositoryImpl{}, userRepo)
				assert.IsType(t, &db.TaskRepositoryImpl{}, taskRepo)
			}),
		)
		app.RequireStart().RequireStop()
	})
}
