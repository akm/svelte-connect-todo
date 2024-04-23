package main

import (
	"net/http"
	"os"
	"strings"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
)

// https://connectrpc.com/docs/cors
// https://github.com/connectrpc/cors-go

// withCORS adds CORS support to a Connect HTTP handler.

var (
	allowedMethods = connectcors.AllowedMethods()
	allowedHeaders = connectcors.AllowedHeaders()
	exposedHeaders = connectcors.ExposedHeaders()
)

func withCORS(connectHandler http.Handler) http.Handler {
	allowedOrigins := strings.Split(os.Getenv("APP_CORS_ALLOW_ORIGINS"), ",")
	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   allowedMethods,
		AllowedHeaders:   allowedHeaders,
		ExposedHeaders:   exposedHeaders,
		AllowCredentials: true,
		MaxAge:           7200, // 2 hours in seconds
	})
	return c.Handler(connectHandler)
}
