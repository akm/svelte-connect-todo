package slog

import (
	"fmt"
	"log/slog"
)

var (
	origDefault    = slog.Default
	origSetDefault = slog.SetDefault
)

var defaultLogger Logger

func Default() Logger {
	if defaultLogger == nil {
		defaultLogger = &loggerImpl{origLogger: origDefault()}
	}
	return defaultLogger
}

func SetDefault(l Logger) error {
	impl, ok := l.(*loggerImpl)
	if !ok {
		return fmt.Errorf("invalid logger type [%T]", l)
	}
	origSetDefault(impl.origLogger)
	defaultLogger = l
	return nil
}
