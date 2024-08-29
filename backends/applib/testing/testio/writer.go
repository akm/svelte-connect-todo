package testio

import (
	"io"
	"testing"
)

type Writer struct {
	t *testing.T
}

var _ io.Writer = (*Writer)(nil)

func New(t *testing.T) *Writer {
	return &Writer{t: t}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.t.Helper()
	w.t.Logf("%s", p)
	return len(p), nil
}
