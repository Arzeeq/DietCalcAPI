package logger

import (
	"dietcalc/internal/config"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type MyLogger struct {
	*slog.Logger
}

func New(env string) *MyLogger {
	var loggerLevel slog.Level

	switch env {
	case config.EnvLocal:
		loggerLevel = slog.LevelDebug
	case config.EnvDev:
		loggerLevel = slog.LevelDebug
	case config.EnvProd:
		loggerLevel = slog.LevelInfo
	}

	return &MyLogger{
		slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: loggerLevel})),
	}
}

func ErrAttr(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func (l *MyLogger) ReplyHTTPError(w http.ResponseWriter, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	errJSONEncode := json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	if errJSONEncode != nil {
		l.Error("failed to write HTTP error", ErrAttr(errJSONEncode))
	}
}
