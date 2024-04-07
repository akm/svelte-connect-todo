package taskservices

import (
	"context"

	v1 "apisvr/gen/task/v1"
	"apisvr/gen/task/v1/taskv1connect"

	"connectrpc.com/connect"
)

type TaskService struct{}

var _ taskv1connect.TaskServiceHandler = (*TaskService)(nil)

type Task = v1.Task

var taskStore = []*Task{
	{Id: 1, Name: "task1", Status: v1.Status_STATUS_DONE},
	{Id: 2, Name: "task2", Status: v1.Status_STATUS_TODO},
}

func (s *TaskService) List(context.Context, *connect.Request[v1.TaskListRequest]) (*connect.Response[v1.TaskListResponse], error) {
	resp := &connect.Response[v1.TaskListResponse]{
		Msg: &v1.TaskListResponse{
			Items: taskStore,
			Total: uint64(len(taskStore)),
		},
	}
	return resp, nil
}

func (s *TaskService) Show(ctx context.Context, req *connect.Request[v1.TaskId]) (*connect.Response[v1.Task], error) {
	for _, task := range taskStore {
		if task.Id == req.Msg.Id {
			return &connect.Response[v1.Task]{Msg: task}, nil
		}
	}
	return nil, nil
}

func (s *TaskService) Create(ctx context.Context, req *connect.Request[v1.TaskCreateRequest]) (*connect.Response[v1.Task], error) {
	var id uint64
	if len(taskStore) == 0 {
		id = 1
	} else {
		id = taskStore[len(taskStore)-1].Id + 1
	}
	task := &Task{
		Id:     id,
		Name:   req.Msg.Name,
		Status: req.Msg.Status,
	}
	taskStore = append(taskStore, task)
	return &connect.Response[v1.Task]{Msg: task}, nil
}

func (s *TaskService) Update(ctx context.Context, req *connect.Request[v1.Task]) (*connect.Response[v1.Task], error) {
	for i, task := range taskStore {
		if task.Id == req.Msg.Id {
			taskStore[i] = req.Msg
			return &connect.Response[v1.Task]{Msg: req.Msg}, nil
		}
	}
	return nil, nil
}

func (s *TaskService) Delete(ctx context.Context, req *connect.Request[v1.TaskId]) (*connect.Response[v1.TaskId], error) {
	for i, task := range taskStore {
		if task.Id == req.Msg.Id {
			taskStore = append(taskStore[:i], taskStore[i+1:]...)
			return &connect.Response[v1.TaskId]{Msg: req.Msg}, nil
		}
	}
	return nil, nil
}
