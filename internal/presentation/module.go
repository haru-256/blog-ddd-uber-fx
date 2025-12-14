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
		// サーバーの生成: routesグループを一括注入
		fx.Annotate(
			server.NewServiceServer,
			fx.ParamTags("", "", `group:"routes"`), // cfg, logger, [routes]
		),
		// Handlerの登録: それぞれ "routes" グループの一員として登録
		fx.Annotate(
			server.NewUserServiceHandler,
			fx.As(new(server.RouteRegistrar)),
			fx.ResultTags(`group:"routes"`),
		),
		fx.Annotate(
			server.NewTaskServiceHandler,
			fx.As(new(server.RouteRegistrar)),
			fx.ResultTags(`group:"routes"`),
		),
	),
	// ライフサイクルフックを登録
	fx.Invoke(server.RegisterLifecycleHooks),
)
