package testslog

import (
	"applib/log/slog"
	"applib/testing/testio"
	"io"
	"os"
	"testing"
)

func Setup(t *testing.T) {
	testLog := os.Getenv("TEST_LOG_OUTPUT")
	if testLog == "" {
		testLog = "stdout"
	}
	var w io.Writer
	switch testLog {
	case "stdout":
		w = os.Stdout
	case "stderr":
		w = os.Stderr
	case "test":
		w = testio.NewWriter(t)
	default:
		t.Fatalf("invalid TEST_LOG_OUTPUT: %s", testLog)
	}
	logger := slog.NewLogger(w, slog.LevelDebug, slog.NewTextHandler)
	slog.SetDefault(logger)
}
