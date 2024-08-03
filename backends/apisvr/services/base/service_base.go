package base

import (
	"context"
	"database/sql"
	"log"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoreflect"
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

func (s *ServiceBase) ValidateMsg(ctx context.Context, msg protoreflect.ProtoMessage) error {
	validator, err := protovalidate.New()
	if err != nil {
		return err
	}
	if err = validator.Validate(msg); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	return nil
}
