package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	sqldblogger "github.com/simukti/sqldb-logger"

	"applib/log/slog"

	"apisvr/sqldb-logger/logadapter/slogadapter"
)

func connectDB(logger slog.Logger) (*sql.DB, error) {
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		return nil, fmt.Errorf("DB_DSN is not set")
	}

	pool, err := sql.Open("mysql", dbDSN)
	if err != nil {
		return nil, err
	}

	origLogger, ok := slog.ToSlogLogger(logger)
	if !ok {
		return nil, fmt.Errorf("logger is not a slog.Logger")
	}

	loggerAdapter := slogadapter.New(origLogger)
	pool = sqldblogger.OpenDriver(dbDSN, pool.Driver(), loggerAdapter)

	return pool, nil
}
