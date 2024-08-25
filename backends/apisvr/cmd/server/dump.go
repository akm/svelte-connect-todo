package main

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
)

func withRequestDumping(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			logger.Error("Error dumping request", slog.String("error", err.Error()))
		} else {
			logger.Debug("Request", slog.String("dump", string(dump)))
		}
		next.ServeHTTP(w, r)
	})
}
