package testlogadapter

import (
	"testing"

	"applib/log/slog"
	"applib/testing/testio"

	sqldblogger "github.com/simukti/sqldb-logger"

	slogadapter "apisvr/sqldb-logger/logadapter/slogadapter"
)

// New creates a log adapter from sqldblogger.Logger to an slog.Logger one.
func New(t *testing.T) sqldblogger.Logger {
	return slogadapter.New(
		slog.NewLogger(testio.NewWriter(t), slog.LevelDebug, slog.NewTextHandler),
	)
}
