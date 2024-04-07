package taskservices

import (
	"context"

	v1 "apisvr/gen/task/v1"
	"apisvr/gen/task/v1/taskv1connect"

	"connectrpc.com/connect"
)

type TaskService struct{}

var _ taskv1connect.TaskServiceHandler = (*TaskService)(nil)

func (s *TaskService) List(context.Context, *connect.Request[v1.TaskListRequest]) (*connect.Response[v1.TaskListResponse], error) {
	return nil, nil
}

func (s *TaskService) Show(context.Context, *connect.Request[v1.TaskId]) (*connect.Response[v1.Task], error) {
	return nil, nil
}

func (s *TaskService) Create(context.Context, *connect.Request[v1.TaskCreateRequest]) (*connect.Response[v1.Task], error) {
	return nil, nil
}

func (s *TaskService) Update(context.Context, *connect.Request[v1.Task]) (*connect.Response[v1.Task], error) {
	return nil, nil
}

func (s *TaskService) Delete(context.Context, *connect.Request[v1.TaskId]) (*connect.Response[v1.TaskId], error) {
	return nil, nil
}
