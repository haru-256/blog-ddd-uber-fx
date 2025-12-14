package presentation

import (
	"github.com/haru-256/blog-ddd-uber-fx/internal/application"
	"github.com/haru-256/blog-ddd-uber-fx/internal/presentation/server"
	"go.uber.org/fx"
)

// Module はプレゼンテーション層のFxモジュールです。
var Module = fx.Module(
	"presentation",
	application.Module,
	fx.Provide(
		fx.Annotate(
			server.NewServiceServerConfig,
			fx.ParamTags(`name:"serverPort"`),
		),
		server.NewServiceServer,
		server.NewUserServiceHandler,
		server.NewTaskServiceHandler,
	),
	// ライフサイクルフックを登録
	fx.Invoke(server.RegisterLifecycleHooks),
)
