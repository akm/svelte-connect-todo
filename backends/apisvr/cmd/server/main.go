package main

import (
	"net/http"

	"connectrpc.com/authn"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"apisvr/gen/task/v1/taskv1connect"
	"apisvr/services/auth"
	taskservices "apisvr/services/task_services"
)

func main() {
	mux := http.NewServeMux()

	// Instantiate the YOUR services and Mount them here.
	authmw := authn.NewMiddleware(auth.Authenticate)

	taskService := &taskservices.TaskService{}
	path, handler := taskv1connect.NewTaskServiceHandler(taskService)
	mux.Handle(path, authmw.Wrap(handler))

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		withCORS(h2c.NewHandler(mux, &http2.Server{})),
	)
}
