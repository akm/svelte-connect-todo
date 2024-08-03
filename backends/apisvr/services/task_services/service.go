package taskservices

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"

	v1 "apisvr/gen/task/v1"
	"apisvr/gen/task/v1/taskv1connect"
	"apisvr/services/base"

	"biz/models"
)

type TaskService struct {
	base.ServiceBase
}

func NewTaskService(pool *sql.DB) *TaskService {
	return &TaskService{ServiceBase: *base.NewServiceBase("TaskService", pool)}
}

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

func (s *TaskService) List(ctx context.Context, req *connect.Request[v1.TaskServiceListRequest]) (*connect.Response[v1.TaskServiceListResponse], error) {
	s.StartAction(ctx, "List")

	queries := models.New(s.Pool)
	dbTasks, err := queries.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*v1.TaskResponse, len(dbTasks))
	for i, task := range dbTasks {
		var st v1.TaskStatus
		switch task.Status {
		case models.TasksStatusTodo:
			st = v1.TaskStatus_TODO
		case models.TasksStatusDone:
			st = v1.TaskStatus_DONE
		default:
			st = v1.TaskStatus_UNKNOWN_UNSPECIFIED
		}

		results[i] = &v1.TaskResponse{
			Id:     task.ID,
			Name:   task.Name,
			Status: st,
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
	s.StartAction(ctx, "Show")

	queries := models.New(s.Pool)
	task, err := queries.GetTask(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	var st v1.TaskStatus
	switch task.Status {
	case models.TasksStatusTodo:
		st = v1.TaskStatus_TODO
	case models.TasksStatusDone:
		st = v1.TaskStatus_DONE
	default:
		st = v1.TaskStatus_UNKNOWN_UNSPECIFIED
	}

	result := &v1.TaskResponse{
		Id:     task.ID,
		Name:   task.Name,
		Status: st,
	}
	return &connect.Response[v1.TaskResponse]{Msg: result}, nil
}

func (s *TaskService) Create(ctx context.Context, req *connect.Request[v1.TaskServiceCreateRequest]) (*connect.Response[v1.TaskResponse], error) {
	s.StartAction(ctx, "Create")

	validator, err := protovalidate.New()
	if err != nil {
		return nil, err
	}
	if err = validator.Validate(req.Msg); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	} else {
		log.Println("validation succeeded")
	}

	queries := models.New(s.Pool)

	var st models.TasksStatus
	switch req.Msg.Status {
	case v1.TaskStatus_TODO:
		st = models.TasksStatusTodo
	case v1.TaskStatus_DONE:
		st = models.TasksStatusDone
	default:
		return nil, fmt.Errorf("unknown status: %v", req.Msg.Status)
	}

	task := models.CreateTaskParams{
		Name:   req.Msg.Name,
		Status: st,
	}
	res, err := queries.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	{
		var st v1.TaskStatus
		switch task.Status {
		case models.TasksStatusTodo:
			st = v1.TaskStatus_TODO
		case models.TasksStatusDone:
			st = v1.TaskStatus_DONE
		default:
			st = v1.TaskStatus_UNKNOWN_UNSPECIFIED
		}

		result := &v1.TaskResponse{
			Id:     uint64(id),
			Name:   task.Name,
			Status: st,
		}
		return &connect.Response[v1.TaskResponse]{Msg: result}, nil
	}
}

func (s *TaskService) Update(ctx context.Context, req *connect.Request[v1.TaskServiceUpdateRequest]) (*connect.Response[v1.TaskResponse], error) {
	s.StartAction(ctx, "Update")

	validator, err := protovalidate.New()
	if err != nil {
		return nil, err
	}
	if err = validator.Validate(req.Msg); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	} else {
		log.Println("validation succeeded")
	}

	queries := models.New(s.Pool)

	var st models.TasksStatus
	switch req.Msg.Status {
	case v1.TaskStatus_TODO:
		st = models.TasksStatusTodo
	case v1.TaskStatus_DONE:
		st = models.TasksStatusDone
	default:
		return nil, fmt.Errorf("unknown status: %v", req.Msg.Status)
	}

	task := models.UpdateTaskParams{
		ID:     req.Msg.Id,
		Name:   req.Msg.Name,
		Status: st,
	}
	if err := queries.UpdateTask(ctx, task); err != nil {
		return nil, err
	}

	{
		var st v1.TaskStatus
		switch task.Status {
		case models.TasksStatusTodo:
			st = v1.TaskStatus_TODO
		case models.TasksStatusDone:
			st = v1.TaskStatus_DONE
		default:
			st = v1.TaskStatus_UNKNOWN_UNSPECIFIED
		}

		result := &v1.TaskResponse{
			Id:     req.Msg.Id,
			Name:   task.Name,
			Status: st,
		}
		return &connect.Response[v1.TaskResponse]{Msg: result}, nil
	}
}

func (s *TaskService) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.TaskResponse], error) {
	s.StartAction(ctx, "Delete")

	queries := models.New(s.Pool)
	task, err := queries.GetTask(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	var st v1.TaskStatus
	switch task.Status {
	case models.TasksStatusTodo:
		st = v1.TaskStatus_TODO
	case models.TasksStatusDone:
		st = v1.TaskStatus_DONE
	default:
		st = v1.TaskStatus_UNKNOWN_UNSPECIFIED
	}

	result := &v1.TaskResponse{
		Id:     task.ID,
		Name:   task.Name,
		Status: st,
	}

	if err := queries.DeleteTask(ctx, req.Msg.Id); err != nil {
		return nil, err
	}

	return &connect.Response[v1.TaskResponse]{Msg: result}, nil
}
