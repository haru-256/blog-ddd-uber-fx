package application_test

import (
	"io"
	"log/slog"
	"testing"

	"github.com/haru-256/blog-ddd-uber-fx/internal/application"
	"github.com/haru-256/blog-ddd-uber-fx/internal/application/service"
	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestModule(t *testing.T) {
	t.Run("Provide Services", func(t *testing.T) {
		app := fxtest.New(t,
			application.Module,
			fx.NopLogger,
			// Suppress logger output
			fx.Decorate(func() (*slog.Logger, error) {
				return slog.New(slog.NewTextHandler(io.Discard, nil)), nil
			}),
			fx.Supply(
				fx.Annotated{Name: "logLevel", Target: "INFO"},
				fx.Annotated{Name: "logFormat", Target: "json"},
				fx.Annotated{Name: "addSource", Target: true},
			),
			fx.Invoke(func(
				userService service.UserService,
				taskService service.TaskService,
				userRepo repositories.UserRepository,
				taskRepo repositories.TaskRepository,
			) {
				assert.NotNil(t, userService)
				assert.NotNil(t, taskService)
				assert.NotNil(t, userRepo)
				assert.NotNil(t, taskRepo)
			}),
		)
		app.RequireStart().RequireStop()
	})
}
