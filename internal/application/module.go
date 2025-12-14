package application

import (
	"github.com/haru-256/blog-ddd-uber-fx/internal/application/service"
	"github.com/haru-256/blog-ddd-uber-fx/internal/infrastructure"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"application",
	infrastructure.Module,
	fx.Provide(
		fx.Annotate(
			service.NewUserServiceImpl,
			fx.As(new(service.UserService)),
		),
		fx.Annotate(
			service.NewTaskServiceImpl,
			fx.As(new(service.TaskService)),
		),
	),
)
