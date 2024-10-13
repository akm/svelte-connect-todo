package slog

type HandlerFunc func(Handler) Handler

type HandlerFuncs []HandlerFunc

func (fns HandlerFuncs) Wrap(h Handler) Handler {
	for i := len(fns) - 1; i >= 0; i-- {
		h = fns[i](h)
	}
	return h
}

var handlerFuncs HandlerFuncs

func RegisterHandlerFunc(f HandlerFunc) {
	handlerFuncs = append(handlerFuncs, f)
}
