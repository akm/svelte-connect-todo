package taskservices

import (
	"context"
	"fmt"
	"log"

	v1 "apisvr/gen/task/v1"
	"apisvr/gen/task/v1/taskv1connect"

	"connectrpc.com/connect"
)

type TaskService struct{}

var _ taskv1connect.TaskServiceHandler = (*TaskService)(nil)

type Task struct {
	Id     uint64
	Name   string
	Status v1.TaskStatus
}

var taskStore = []*Task{
	{Id: 1, Name: "task1", Status: v1.TaskStatus_DONE},
	{Id: 2, Name: "task2", Status: v1.TaskStatus_TODO},
}

func (s *TaskService) List(context.Context, *connect.Request[v1.TaskServiceListRequest]) (*connect.Response[v1.TaskServiceListResponse], error) {
	log.Printf("TaskService.List")

	results := make([]*v1.TaskResponse, len(taskStore))
	for i, task := range taskStore {
		results[i] = &v1.TaskResponse{
			Id:     task.Id,
			Name:   task.Name,
			Status: task.Status,
		}
	}

	resp := &connect.Response[v1.TaskServiceListResponse]{
		Msg: &v1.TaskServiceListResponse{
			Items: results,
			Total: uint64(len(results)),
		},
	}
	return resp, nil
}

func (s *TaskService) Show(ctx context.Context, req *connect.Request[v1.ShowRequest]) (*connect.Response[v1.TaskResponse], error) {
	log.Printf("TaskService.Show")

	for _, task := range taskStore {
		if task.Id == req.Msg.Id {
			result := &v1.TaskResponse{
				Id:     task.Id,
				Name:   task.Name,
				Status: task.Status,
			}
			return &connect.Response[v1.TaskResponse]{Msg: result}, nil
		}
	}

	return nil, fmt.Errorf("task not found")
}

func (s *TaskService) Create(ctx context.Context, req *connect.Request[v1.TaskServiceCreateRequest]) (*connect.Response[v1.TaskResponse], error) {
	log.Printf("TaskService.Create")
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

	result := &v1.TaskResponse{
		Id:     task.Id,
		Name:   task.Name,
		Status: task.Status,
	}
	return &connect.Response[v1.TaskResponse]{Msg: result}, nil
}

func (s *TaskService) Update(ctx context.Context, req *connect.Request[v1.TaskServiceUpdateRequest]) (*connect.Response[v1.TaskResponse], error) {
	log.Printf("TaskService.Update")
	for i, task := range taskStore {
		if task.Id == req.Msg.Id {
			task := &Task{
				Id:     req.Msg.Id,
				Name:   req.Msg.Name,
				Status: req.Msg.Status,
			}
			taskStore[i] = task

			result := &v1.TaskResponse{
				Id:     task.Id,
				Name:   task.Name,
				Status: task.Status,
			}
			return &connect.Response[v1.TaskResponse]{Msg: result}, nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func (s *TaskService) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.TaskResponse], error) {
	log.Printf("TaskService.Delete")
	for i, task := range taskStore {
		if task.Id == req.Msg.Id {
			task := taskStore[i]
			taskStore = append(taskStore[:i], taskStore[i+1:]...)

			result := &v1.TaskResponse{
				Id:     task.Id,
				Name:   task.Name,
				Status: task.Status,
			}
			return &connect.Response[v1.TaskResponse]{Msg: result}, nil
		}
	}
	return nil, fmt.Errorf("task not found")
}
