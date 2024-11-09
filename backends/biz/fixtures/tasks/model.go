package tasks

import "biz/models"

type Task struct {
	models.Task
}

type Option = func(*Task)

func Name(v string) Option               { return func(t *Task) { t.Name = v } }
func Status(v models.TasksStatus) Option { return func(t *Task) { t.Status = v } }
