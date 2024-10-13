package slog

import (
	"io"
	"log/slog"
	orig "log/slog"
	"os"
	"strings"
)

func New(w io.Writer) (Logger, error) {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = "INFO"
	}
	var level Level
	if err := level.UnmarshalText([]byte(logLevelStr)); err != nil {
		return nil, err
	}

	var newHandler func(w io.Writer, opts *orig.HandlerOptions) Handler
	switch strings.ToLower(os.Getenv("LOG_FORMAT")) {
	case "text":
		newHandler = NewTextHandler
	default:
		newHandler = NewJSONHandler
	}

	return NewLogger(w, level, newHandler), nil
}

func NewLogger(w io.Writer, level Level, newHandler func(w io.Writer, opts *orig.HandlerOptions) Handler) Logger {
	opts := &slog.HandlerOptions{Level: level}
	handler := handlerFuncs.Wrap(newHandler(w, opts))
	return &loggerImpl{origLogger: slog.New(handler)}
}
