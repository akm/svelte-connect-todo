package slog

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

func New(w io.Writer) (Logger, error) {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = "INFO"
	}
	var level slog.Level
	if err := level.UnmarshalText([]byte(logLevelStr)); err != nil {
		return nil, err
	}
	opts := &slog.HandlerOptions{Level: level}

	var handler slog.Handler
	switch strings.ToLower(os.Getenv("LOG_FORMAT")) {
	case "text":
		handler = slog.NewTextHandler(w, opts)
	default:
		handler = slog.NewJSONHandler(w, opts)
	}

	return &loggerImpl{origLogger: slog.New(handler)}, nil
}
