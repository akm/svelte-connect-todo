package main

import (
	"database/sql"
	"fmt"
	"os"

	sqldbloggerslog "github.com/akm/sqldb-logger-slog"
	_ "github.com/go-sql-driver/mysql"
	sqldblogger "github.com/simukti/sqldb-logger"

	"applib/log/slog"
)

func connectDB(logger *slog.Logger) (*sql.DB, error) {
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		return nil, fmt.Errorf("DB_DSN is not set")
	}

	pool, err := sql.Open("mysql", dbDSN)
	if err != nil {
		return nil, err
	}

	loggerAdapter := sqldbloggerslog.New(logger)
	pool = sqldblogger.OpenDriver(dbDSN, pool.Driver(), loggerAdapter)

	return pool, nil
}
