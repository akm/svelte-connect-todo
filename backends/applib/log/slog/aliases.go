package slog

import (
	"io"
	orig "log/slog"
)

type (
	Level = orig.Level
	Attr  = orig.Attr

	Handler        = orig.Handler
	HandlerOptions = orig.HandlerOptions

	origLogger = orig.Logger
)

var (
	String   = orig.String
	Int64    = orig.Int64
	Int      = orig.Int
	Uint64   = orig.Uint64
	Float64  = orig.Float64
	Bool     = orig.Bool
	Time     = orig.Time
	Duration = orig.Duration
	Group    = orig.Group
	Any      = orig.Any

	NewJSONHandlerOrig = orig.NewJSONHandler
	NewTextHandlerOrig = orig.NewTextHandler

	NewJSONHandler = func(w io.Writer, opts *HandlerOptions) Handler { return NewJSONHandlerOrig(w, opts) }
	NewTextHandler = func(w io.Writer, opts *HandlerOptions) Handler { return NewTextHandlerOrig(w, opts) }
)

const (
	LevelError = orig.LevelError
	LevelWarn  = orig.LevelWarn
	LevelInfo  = orig.LevelInfo
	LevelDebug = orig.LevelDebug
)
