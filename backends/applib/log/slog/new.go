package slog

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/akm/slogw"
)

var ErrInvalidLogFormat = fmt.Errorf("invalid LOG_FORMAT. Use 'text' or 'json'")

func New(w io.Writer) (*Logger, error) {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = "INFO"
	}
	var level Level
	if err := level.UnmarshalText([]byte(logLevelStr)); err != nil {
		return nil, err
	}

	var newHandler func(w io.Writer, opts *HandlerOptions) Handler
	logFormat := strings.ToLower(os.Getenv("LOG_FORMAT"))
	if logFormat == "" {
		logFormat = "json"
	}
	switch logFormat {
	case "text":
		newHandler = NewTextHandler
	case "json":
		newHandler = NewJSONHandler
	default:
		return nil, ErrInvalidLogFormat
	}

	return NewLogger(w, level, newHandler), nil
}

func NewLogger(w io.Writer, level Level, newHandler func(w io.Writer, opts *HandlerOptions) Handler) *Logger {
	opts := &HandlerOptions{Level: level}
	return slogw.New(newHandler(w, opts))
}
