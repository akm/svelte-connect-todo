package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func withRequestDumping(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Printf("Error dumping request: %v", err)
		} else {
			log.Printf("Request: %s", dump)
		}
		next.ServeHTTP(w, r)
	})
}
