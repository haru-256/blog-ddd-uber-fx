package infrastructure

import (
	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/repositories"
	"github.com/haru-256/blog-ddd-uber-fx/internal/infrastructure/db"
	"github.com/haru-256/blog-ddd-uber-fx/internal/infrastructure/logger"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure",
	fx.Provide(
		fx.Annotate(
			logger.NewLogger,
			fx.ParamTags(`name:"logLevel"`, `name:"logFormat"`, `name:"addSource"`),
		),
		fx.Annotate(
			db.NewUserRepositoryImpl,
			fx.As(new(repositories.UserRepository)),
		),
		fx.Annotate(
			db.NewTaskRepositoryImpl,
			fx.As(new(repositories.TaskRepository)),
		),
	),
	// Privateにすることで他のモジュールからは見えない
	fx.Provide(
		db.NewDatabase,
		fx.Private,
	),
)
