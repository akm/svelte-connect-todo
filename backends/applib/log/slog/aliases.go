package slog

import (
	orig "log/slog"
)

type (
	Attr    = orig.Attr
	Handler = orig.Handler
	Level   = orig.Level

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
)
