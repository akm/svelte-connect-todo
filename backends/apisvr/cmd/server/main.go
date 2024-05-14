package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/authn"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"apisvr/gen/task/v1/taskv1connect"
	"apisvr/services/auth"
	taskservices "apisvr/services/task_services"
)

func main() {
	log.Printf("Starting api server")

	mux := http.NewServeMux()

	// Instantiate the YOUR services and Mount them here.
	authmw := authn.NewMiddleware(auth.Authenticate)

	taskService := &taskservices.TaskService{}
	path, handler := taskv1connect.NewTaskServiceHandler(taskService)
	mux.Handle(path, authmw.Wrap(handler))

	// https://cloud.google.com/run/docs/triggering/grpc?hl=ja
	serverHostAndPort := os.Getenv("APP_SERVER_HOST_AND_PORT")
	if serverHostAndPort == "" {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		serverHostAndPort = ":" + port
	}

	// https://connectrpc.com/docs/go/deployment/
	// https://github.com/connectrpc/examples-go/blob/main/cmd/demoserver/main.go
	muxHandler := withCORS(h2c.NewHandler(mux, &http2.Server{}))
	muxHandler = withRequestDumping(muxHandler)

	srv := &http.Server{
		Addr:              serverHostAndPort,
		Handler:           muxHandler,
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	log.Printf("Starting api server on %s", serverHostAndPort)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP listen and serve: %v", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err) //nolint:gocritic
	}
}
