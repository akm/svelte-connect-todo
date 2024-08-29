package taskservices

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	taskv1 "apisvr/gen/task/v1"

	"applib/sqldb-logger/logadapter/slogadapter/testslogadapter"

	"connectrpc.com/connect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

var (
	pool     *sql.DB
	fixtures *testfixtures.Loader
)

var dsn = os.Getenv("DB_DSN")

func TestMain(m *testing.M) {
	var err error

	// Open connection to the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be deleted.
	log.Printf("DB_DSN: %s", dsn)
	pool, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}

	fixtureDir := os.Getenv("TEST_PATH_TO_FIXTURES")
	log.Printf("TEST_PATH_TO_FIXTURES: %s", fixtureDir)
	fixtures, err = testfixtures.New(
		testfixtures.Database(pool),        // You database connection
		testfixtures.Dialect("mysql"),      // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory(fixtureDir), // The directory containing the YAML files
		testfixtures.SkipResetSequences(),  // Disable the execution of the SQL command that resets the sequences
	)
	if err != nil {
		log.Fatalf("unable to load fixtures: %v", err)
	}

	os.Exit(m.Run())
}

func prepareTestDatabase(t *testing.T) {
	err := fixtures.Load()
	assert.NoError(t, err)
}

func TestTaskServiceList(t *testing.T) {
	prepareTestDatabase(t)

	adapter := testslogadapter.New(t)
	pool := sqldblogger.OpenDriver(dsn, pool.Driver(), adapter)

	ctx := context.Background()

	srv := NewTaskService(pool)
	resp, err := srv.List(ctx, &connect.Request[taskv1.TaskServiceListRequest]{
		Msg: &taskv1.TaskServiceListRequest{},
	})
	assert.NoError(t, err)

	assert.Equal(t, "Survey the market", resp.Msg.Items[0].Name)
	assert.Equal(t, taskv1.TaskStatus_DONE, resp.Msg.Items[0].Status)

	assert.Equal(t, "Plan the project", resp.Msg.Items[1].Name)
	assert.Equal(t, taskv1.TaskStatus_TODO, resp.Msg.Items[1].Status)
}

func TestTaskServiceShow(t *testing.T) {
	prepareTestDatabase(t)

	adapter := testslogadapter.New(t)
	pool := sqldblogger.OpenDriver(dsn, pool.Driver(), adapter)

	srv := NewTaskService(pool)

	t.Run("valid id", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Show(ctx, &connect.Request[taskv1.ShowRequest]{
			Msg: &taskv1.ShowRequest{Id: 1},
		})
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Msg.Id)
		assert.Equal(t, "Survey the market", resp.Msg.Name)
		assert.Equal(t, taskv1.TaskStatus_DONE, resp.Msg.Status)
	})
	t.Run("invalid id", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Show(ctx, &connect.Request[taskv1.ShowRequest]{
			Msg: &taskv1.ShowRequest{Id: 999999},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
		assert.Nil(t, resp)
	})
}

func TestTaskServiceCreate(t *testing.T) {
	prepareTestDatabase(t)

	adapter := testslogadapter.New(t)
	pool := sqldblogger.OpenDriver(dsn, pool.Driver(), adapter)

	srv := NewTaskService(pool)

	t.Run("valid task", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Create(ctx, &connect.Request[taskv1.TaskServiceCreateRequest]{
			Msg: &taskv1.TaskServiceCreateRequest{
				Name:   "Implement the project",
				Status: taskv1.TaskStatus_TODO,
			},
		})
		assert.NoError(t, err)
		assert.Greater(t, resp.Msg.Id, uint64(2))
		assert.Equal(t, "Implement the project", resp.Msg.Name)
		assert.Equal(t, taskv1.TaskStatus_TODO, resp.Msg.Status)
	})
	t.Run("empty name", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Create(ctx, &connect.Request[taskv1.TaskServiceCreateRequest]{
			Msg: &taskv1.TaskServiceCreateRequest{
				Name:   "",
				Status: taskv1.TaskStatus_TODO,
			},
		})
		if assert.Error(t, err) {
			if assert.IsTypef(t, &connect.Error{}, err, "expected *connect.Error, got %T", err) {
				details := err.(*connect.Error).Details()
				if assert.Len(t, details, 1) {
					detail := details[0]
					assert.Equal(t, "google.rpc.BadRequest.FieldViolation", detail.Type())
					val, err := detail.Value()
					assert.NoError(t, err)
					if assert.IsType(t, &errdetails.BadRequest_FieldViolation{}, val, "expected *errdetails.BadRequest_FieldViolation, got %T", val) {
						fieldViolation := val.(*errdetails.BadRequest_FieldViolation)
						assert.Equal(t, "name", fieldViolation.Field)
						assert.Equal(t, "value length must be at least 1 characters", fieldViolation.Description)
					}
				}
			}
		}
		assert.Nil(t, resp)
	})
	t.Run("invalid status", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Create(ctx, &connect.Request[taskv1.TaskServiceCreateRequest]{
			Msg: &taskv1.TaskServiceCreateRequest{
				Name:   "Implement the project",
				Status: taskv1.TaskStatus_UNKNOWN_UNSPECIFIED,
			},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unknown status")
		assert.Nil(t, resp)
	})
}

func TestTaskServiceUpdate(t *testing.T) {
	prepareTestDatabase(t)

	adapter := testslogadapter.New(t)
	pool := sqldblogger.OpenDriver(dsn, pool.Driver(), adapter)

	srv := NewTaskService(pool)

	t.Run("valid task", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Update(ctx, &connect.Request[taskv1.TaskServiceUpdateRequest]{
			Msg: &taskv1.TaskServiceUpdateRequest{
				Id:     2,
				Name:   "Make the project schedule",
				Status: taskv1.TaskStatus_DONE,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, uint64(2), resp.Msg.Id)
		assert.Equal(t, "Make the project schedule", resp.Msg.Name)
		assert.Equal(t, taskv1.TaskStatus_DONE, resp.Msg.Status)
	})
	t.Run("empty name", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Update(ctx, &connect.Request[taskv1.TaskServiceUpdateRequest]{
			Msg: &taskv1.TaskServiceUpdateRequest{
				Id:     2,
				Name:   "",
				Status: taskv1.TaskStatus_DONE,
			},
		})
		if assert.Error(t, err) {
			if assert.IsTypef(t, &connect.Error{}, err, "expected *connect.Error, got %T", err) {
				details := err.(*connect.Error).Details()
				if assert.Len(t, details, 1) {
					detail := details[0]
					assert.Equal(t, "google.rpc.BadRequest.FieldViolation", detail.Type())
					val, err := detail.Value()
					assert.NoError(t, err)
					if assert.IsType(t, &errdetails.BadRequest_FieldViolation{}, val, "expected *errdetails.BadRequest_FieldViolation, got %T", val) {
						fieldViolation := val.(*errdetails.BadRequest_FieldViolation)
						assert.Equal(t, "name", fieldViolation.Field)
						assert.Equal(t, "value length must be at least 1 characters", fieldViolation.Description)
					}
				}
			}
		}
		assert.Nil(t, resp)
	})
	t.Run("invalid status", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Update(ctx, &connect.Request[taskv1.TaskServiceUpdateRequest]{
			Msg: &taskv1.TaskServiceUpdateRequest{
				Id:     2,
				Name:   "Make the project schedule",
				Status: taskv1.TaskStatus_UNKNOWN_UNSPECIFIED,
			},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unknown status")
		assert.Nil(t, resp)
	})
	t.Run("invalid id", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Update(ctx, &connect.Request[taskv1.TaskServiceUpdateRequest]{
			Msg: &taskv1.TaskServiceUpdateRequest{
				Id:     999999,
				Name:   "Make the project schedule",
				Status: taskv1.TaskStatus_DONE,
			},
		})
		if assert.IsTypef(t, &connect.Error{}, err, "expected connect.Error, got %T", err) {
			connectErr := err.(*connect.Error)
			assert.Equal(t, connect.CodeNotFound, connectErr.Code())
			assert.Contains(t, connectErr.Message(), "task not found")
		}
		assert.Nil(t, resp)
	})
}

func TestTaskServiceDelete(t *testing.T) {
	prepareTestDatabase(t)

	adapter := testslogadapter.New(t)
	pool := sqldblogger.OpenDriver(dsn, pool.Driver(), adapter)

	srv := NewTaskService(pool)

	t.Run("valid task", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Delete(ctx, &connect.Request[taskv1.DeleteRequest]{
			Msg: &taskv1.DeleteRequest{Id: 1},
		})
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), resp.Msg.Id)
		assert.Equal(t, "Survey the market", resp.Msg.Name)
		assert.Equal(t, taskv1.TaskStatus_DONE, resp.Msg.Status)
	})
	t.Run("invalid id", func(t *testing.T) {
		ctx := context.Background()
		resp, err := srv.Delete(ctx, &connect.Request[taskv1.DeleteRequest]{
			Msg: &taskv1.DeleteRequest{Id: 999999},
		})
		assert.Error(t, err)
		if assert.IsTypef(t, &connect.Error{}, err, "expected connect.Error, got %T", err) {
			connectErr := err.(*connect.Error)
			assert.Equal(t, connect.CodeNotFound, connectErr.Code())
			assert.Contains(t, connectErr.Message(), "task not found")
		}
		assert.Nil(t, resp)
	})
}
