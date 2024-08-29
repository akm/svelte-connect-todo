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
