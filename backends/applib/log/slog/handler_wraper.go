package slog

import (
	"context"
	"log/slog"
)

type HandleFunc = func(context.Context, slog.Record) error

type handlerFuncWrapper struct {
	Handler
	handle HandleFunc
}

var _ Handler = (*handlerFuncWrapper)(nil)

func NewFuncHandlerWrapper(fn func(orig HandleFunc) HandleFunc) HandlerFunc {
	return func(h Handler) Handler {
		handle := fn(h.Handle)
		return &handlerFuncWrapper{Handler: h, handle: handle}
	}
}

func (h *handlerFuncWrapper) Handle(ctx context.Context, rec slog.Record) error {
	return h.handle(ctx, rec)
}
