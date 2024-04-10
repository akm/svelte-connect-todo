package main

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
)

// https://connectrpc.com/docs/cors
// https://github.com/connectrpc/cors-go

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(connectHandler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://localhost:4173"}, // replace with your domain
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200, // 2 hours in seconds
	})
	return c.Handler(connectHandler)
}
