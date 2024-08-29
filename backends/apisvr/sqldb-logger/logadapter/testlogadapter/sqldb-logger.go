package testlogadapter

// Copied from https://github.com/simukti/sqldb-logger/pull/76

import (
	"context"
	"log/slog"
	"testing"

	sqldblogger "github.com/simukti/sqldb-logger"
)

type slogAdapter struct {
	t *testing.T
}

// New creates a log adapter from sqldblogger.Logger to an slog.Logger one.
func New(t *testing.T) sqldblogger.Logger {
	return &slogAdapter{t: t}
}

// Log implement sqldblogger.Logger and converts its levels to corresponding
// log/slog ones.
func (a *slogAdapter) Log(ctx context.Context, sqldbLevel sqldblogger.Level, msg string, data map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(data))
	for k, v := range data {
		attrs = append(attrs, slog.Any(k, v))
	}

	var level slog.Level
	switch sqldbLevel {
	case sqldblogger.LevelError:
		level = slog.LevelError
	case sqldblogger.LevelInfo:
		level = slog.LevelInfo
	case sqldblogger.LevelDebug:
		level = slog.LevelDebug
	default:
		// trace will use slog debug
		level = slog.LevelDebug
	}

	a.t.Logf("[%s] %s %v", level, msg, attrs)
}
