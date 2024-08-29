package testslog

import (
	"applib/log/slog"
	"applib/testing/testio"
	"testing"
)

func New(t *testing.T) slog.Logger {
	return slog.NewLogger(testio.NewWriter(t), slog.LevelDebug, slog.NewTextHandler)
}
