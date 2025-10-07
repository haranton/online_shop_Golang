package logger

import (
	"log/slog"
	"os"
)

func GetLogger(env string) *slog.Logger {

	var logger *slog.Logger
	switch env {
	case "DEBUG":
		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		logger = slog.New(handler)
	case "PRODUCTION":
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		logger = slog.New(handler)
	default:
		logger = slog.Default()
	}

	return logger

}
