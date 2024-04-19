package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"apisvr/gen/session/v1/sessionv1connect"
	"apisvr/gen/task/v1/taskv1connect"
	sessionservice "apisvr/services/session_service"
	taskservices "apisvr/services/task_services"
)

func main() {
	mux := http.NewServeMux()

	// Instantiate the YOUR services and Mount them here.
	{
		svc := new(sessionservice.SessionService)
		path, handler := sessionv1connect.NewSessionServiceHandler(svc)
		mux.Handle(path, handler)
	}
	{
		svc := &taskservices.TaskService{}
		path, handler := taskv1connect.NewTaskServiceHandler(svc)
		mux.Handle(path, handler)
	}

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		withCORS(h2c.NewHandler(mux, &http2.Server{})),
	)
}
