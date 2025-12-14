package main

import (
	"log/slog"

	"github.com/haru-256/blog-ddd-uber-fx/internal/presentation"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	app := fx.New(
		fx.Supply(
			fx.Annotated{Name: "serverPort", Target: "8080"},
			fx.Annotated{Name: "logLevel", Target: "info"},
			fx.Annotated{Name: "logFormat", Target: "json"},
			fx.Annotated{Name: "addSource", Target: false},
		),
		presentation.Module,
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: logger}
		}),
	)
	app.Run()
}
