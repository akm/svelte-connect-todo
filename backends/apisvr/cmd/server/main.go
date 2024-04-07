package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// import (
// 	"apisvr/gen/task/v1/taskv1connect"
// 	taskservices "apisvr/services/task_services"
// )

func main() {
	mux := http.NewServeMux()

	// Instantiate the YOUR services and Mount them here.

	// EXAMPLE:
	// taskService := &taskservices.TaskService{}
	// path, handler := taskv1connect.NewTaskServiceHandler(taskService)
	// mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
