package slog

import (
	"io"
	"os"
	"strings"

	"github.com/akm/slogwrap"
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

	var newHandler func(w io.Writer, opts *HandlerOptions) Handler
	switch strings.ToLower(os.Getenv("LOG_FORMAT")) {
	case "text":
		newHandler = NewTextHandler
	default:
		newHandler = NewJSONHandler
	}

	return NewLogger(w, level, newHandler), nil
}

func NewLogger(w io.Writer, level Level, newHandler func(w io.Writer, opts *HandlerOptions) Handler) Logger {
	opts := &HandlerOptions{Level: level}
	return &loggerImpl{origLogger: slogwrap.New(newHandler(w, opts))}
}
