package main

import (
	"context"
	"errors"
	"fmt"
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
		fmt.Printf("Logger error: %+v", err)
		os.Exit(1)
		return
	}
	if err := slog.SetDefault(logger); err != nil {
		panic(err)
	}

	pool, err := connectDB(logger)
	if err != nil {
		logger.Error("DB connection error", "cause", err)
	}
	defer pool.Close()

	serviceMux := http.NewServeMux()

	// Instantiate the YOUR services and Mount them here.
	authmw := authn.NewMiddleware(auth.Authenticate(logger))

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

	logger.Info("Starting server", "host", serverHostAndPort)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("HTTP listen and serve", "cause", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("HTTP shutdown", "cause", err)
	}
}
