package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func NewLogger(logLevel string, logFormat string, addSource bool) (*slog.Logger, error) {
	// ログレベルを設定
	level := slog.LevelInfo
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
	}

	// ハンドラの設定
	var handler slog.Handler
	switch logFormat {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case "text":
		handler = slog.NewTextHandler(os.Stdout, opts)
	default:
		return nil, fmt.Errorf("invalid log format: %s", logFormat)
	}

	// ロガーを作成して返す
	logger := slog.New(handler)
	return logger, nil
}
