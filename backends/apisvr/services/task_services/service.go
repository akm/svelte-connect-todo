package taskservices

import (
	"context"
	"database/sql"
	"fmt"

	"connectrpc.com/connect"

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

func (s *TaskService) List(ctx context.Context, req *connect.Request[v1.TaskServiceListRequest]) (resp *connect.Response[v1.TaskServiceListResponse], rerr error) {
	rerr = s.Action(ctx, "List", func(ctx context.Context) error {
		queries := models.New(s.Pool)
		dbTasks, err := queries.ListTasks(ctx)
		if err != nil {
			return err
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

		resp = &connect.Response[v1.TaskServiceListResponse]{
			Msg: &v1.TaskServiceListResponse{
				Items: results,
				Total: uint64(len(results)),
			},
		}
		return nil
	})
	return
}

func (s *TaskService) Show(ctx context.Context, req *connect.Request[v1.ShowRequest]) (resp *connect.Response[v1.TaskResponse], rerr error) {
	rerr = s.Action(ctx, "Show", func(ctx context.Context) error {
		queries := models.New(s.Pool)
		task, err := queries.GetTask(ctx, req.Msg.Id)
		if err != nil {
			return s.ToConnectError(err)
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
		resp = &connect.Response[v1.TaskResponse]{Msg: result}
		return nil
	})
	return
}

func (s *TaskService) Create(ctx context.Context, req *connect.Request[v1.TaskServiceCreateRequest]) (resp *connect.Response[v1.TaskResponse], rerr error) {
	rerr = s.Action(ctx, "Create", func(ctx context.Context) error {
		if err := s.ValidateMsg(ctx, req.Msg); err != nil {
			return err
		}

		queries := models.New(s.Pool)

		var st models.TasksStatus
		switch req.Msg.Status {
		case v1.TaskStatus_TODO:
			st = models.TasksStatusTodo
		case v1.TaskStatus_DONE:
			st = models.TasksStatusDone
		default:
			return fmt.Errorf("unknown status: %v", req.Msg.Status)
		}

		task := models.CreateTaskParams{
			Name:   req.Msg.Name,
			Status: st,
		}
		res, err := queries.CreateTask(ctx, task)
		if err != nil {
			return err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err
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
			resp = &connect.Response[v1.TaskResponse]{Msg: result}
			return nil
		}
	})
	return
}

func (s *TaskService) Update(ctx context.Context, req *connect.Request[v1.TaskServiceUpdateRequest]) (resp *connect.Response[v1.TaskResponse], rerr error) {
	rerr = s.Action(ctx, "Update", func(ctx context.Context) error {
		if err := s.ValidateMsg(ctx, req.Msg); err != nil {
			return err
		}

		return s.Transaction(ctx, func(tx *sql.Tx) error {
			qtx := models.New(s.Pool).WithTx(tx)
			if _, err := qtx.GetTaskForUpdate(ctx, req.Msg.Id); err != nil {
				return s.ToConnectError(err)
			}

			var st models.TasksStatus
			switch req.Msg.Status {
			case v1.TaskStatus_TODO:
				st = models.TasksStatusTodo
			case v1.TaskStatus_DONE:
				st = models.TasksStatusDone
			default:
				return fmt.Errorf("unknown status: %v", req.Msg.Status)
			}

			task := models.UpdateTaskParams{
				ID:     req.Msg.Id,
				Name:   req.Msg.Name,
				Status: st,
			}
			if err := qtx.UpdateTask(ctx, task); err != nil {
				return s.ToConnectError(err)
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
				resp = &connect.Response[v1.TaskResponse]{Msg: result}
			}
			return nil
		})
	})
	return
}

func (s *TaskService) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (resp *connect.Response[v1.TaskResponse], rerr error) {
	rerr = s.Action(ctx, "Delete", func(ctx context.Context) error {
		queries := models.New(s.Pool)
		task, err := queries.GetTask(ctx, req.Msg.Id)
		if err != nil {
			return s.ToConnectError(err)
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
			return err
		}

		resp = &connect.Response[v1.TaskResponse]{Msg: result}
		return nil
	})
	return
}
