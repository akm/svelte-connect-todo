package base

import (
	"context"
	"database/sql"
	"log"
)

type ServiceBase struct {
	Name string
	Pool *sql.DB
}

func NewServiceBase(name string, pool *sql.DB) *ServiceBase {
	return &ServiceBase{Name: name, Pool: pool}
}

func (s *ServiceBase) StartAction(ctx context.Context, method string) {
	log.Printf("%s.%s\n", s.Name, method)
}
