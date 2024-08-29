package slog

type SlogLoggerOwner interface {
	SlogLogger() *origLogger
}

func ToSlogLogger(x any) (*origLogger, bool) {
	if x, ok := x.(SlogLoggerOwner); ok {
		return x.SlogLogger(), true
	}
	return nil, false
}
