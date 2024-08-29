package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
		switch verr := err.(type) {
		case *protovalidate.ValidationError:
			// https://connectrpc.com/docs/go/errors
			result := connect.NewError(connect.CodeInvalidArgument, errors.New("validation error"))
			for _, violation := range verr.Violations {
				fieldErr := &errdetails.BadRequest_FieldViolation{
					Field:       violation.GetFieldPath(),
					Description: violation.GetMessage(),
				}
				detailErr, err := connect.NewErrorDetail(fieldErr)
				if err != nil {
					return connect.NewError(connect.CodeInvalidArgument, err)
				}
				result.AddDetail(detailErr)
			}
			return result
		default:
			return connect.NewError(connect.CodeInvalidArgument, err)
		}
	}
	return nil
}

func (s *ServiceBase) ToConnectError(err error) *connect.Error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*connect.Error); ok {
		return err.(*connect.Error)
	}

	switch err {
	case sql.ErrNoRows:
		return connect.NewError(connect.CodeNotFound, fmt.Errorf("task not found"))
	default:
		return connect.NewError(connect.CodeInternal, err)
	}
}

func (s *ServiceBase) Transaction(ctx context.Context, f func(*sql.Tx) error) error {
	// 参考 https://docs.sqlc.dev/en/stable/howto/transactions.html
	tx, err := s.Pool.Begin()
	if err != nil {
		return err
	}

	txErr := f(tx)
	if txErr != nil {
		if err := tx.Rollback(); err != nil {
			slog.Error("failed to rollback transaction", "error", err)
		}
		return txErr
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
