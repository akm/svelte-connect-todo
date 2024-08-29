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

	"applib/log/slog"

	"apisvr/gen/task/v1/taskv1connect"
	"apisvr/services/auth"
	"apisvr/services/images"
	taskservices "apisvr/services/task_services"
)

func main() {
	logger, err := slog.New(os.Stdout)
	if err != nil {
		log.Fatalf("Logger error: %+v", err) //nolint:gocritic
	}

	pool, err := connectDB(logger)
	if err != nil {
		log.Fatalf("DB connection error: %+v", err) //nolint:gocritic
	}
	defer pool.Close()

	serviceMux := http.NewServeMux()

	// Instantiate the YOUR services and Mount them here.
	authmw := authn.NewMiddleware(auth.Authenticate)

	taskService := taskservices.NewTaskService(pool)
	path, handler := taskv1connect.NewTaskServiceHandler(taskService)
	serviceMux.Handle(path, authmw.Wrap(handler))

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
	rootMux := http.NewServeMux()
	rootMux.Handle("GET /images/{id}", http.HandlerFunc(images.GetImage))
	rootMux.Handle("/", h2c.NewHandler(serviceMux, &http2.Server{}))

	serviceMuxHandler := withCORS(rootMux)
	serviceMuxHandler = withRequestDumping(serviceMuxHandler, logger)

	srv := &http.Server{
		Addr:              serverHostAndPort,
		Handler:           serviceMuxHandler,
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	log.Printf("Starting server on %s", serverHostAndPort)

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
