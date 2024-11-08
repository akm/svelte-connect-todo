package tasks

import (
	"biz/models"

	fixtures "github.com/akm/go-fixtures"
)

type Fixtures struct {
	*fixtures.Fixtures[Task]
}

var (
	_ (fixtures.Factory[Task]) = (*Fixtures)(nil)
	_ (fixtures.Getter[Task])  = (*Fixtures)(nil)
)

func NewFixtures(opts ...func(*Fixtures)) *Fixtures {
	r := &Fixtures{}
	for _, opt := range opts {
		opt(r)
	}
	r.Fixtures = fixtures.NewFixtures[Task](r)
	return r
}

func (f *Fixtures) NewSurveyTheMarket(opts ...func(*Task)) *Task {
	return fixtures.NewWithDefaults[Task](opts,
		Name("Survey the market"),
		Status(models.TasksStatusDone),
	)
}

func (f *Fixtures) NewPlanTheProject(opts ...func(*Task)) *Task {
	return fixtures.NewWithDefaults[Task](opts,
		Name("Plan the project"),
		Status(models.TasksStatusTodo),
	)
}

func (f *Fixtures) SurveyTheMarket(opts ...Option) *Task { return f.Get("SurveyTheMarket", opts...) }
func (f *Fixtures) PlanTheProject(opts ...Option) *Task  { return f.Get("PlanTheProject", opts...) }
