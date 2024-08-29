package slog

type loggerImpl struct {
	*origLogger
}

var _ Logger = (*loggerImpl)(nil)

func (x *loggerImpl) With(args ...any) Logger {
	resultImpl := x.origLogger.With(args...)
	return &loggerImpl{origLogger: resultImpl}
}

func (x *loggerImpl) WithGroup(name string) Logger {
	resultImpl := x.origLogger.WithGroup(name)
	return &loggerImpl{origLogger: resultImpl}
}

func (x *loggerImpl) SlogLogger() *origLogger {
	return x.origLogger
}
