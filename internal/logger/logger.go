package logger

import (
	"dietcalc/internal/config"
	"log/slog"
	"os"
)

func Setup(env string) *slog.Logger {
	var loggerLevel slog.Level

	switch env {
	case config.EnvLocal:
		loggerLevel = slog.LevelDebug
	case config.EnvDev:
		loggerLevel = slog.LevelDebug
	case config.EnvProd:
		loggerLevel = slog.LevelInfo
	}

	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: loggerLevel}))
}
