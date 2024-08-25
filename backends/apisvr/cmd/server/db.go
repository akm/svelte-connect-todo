package main

import (
	"apisvr/sqldb-logger/logadapter/slogadapter"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	sqldblogger "github.com/simukti/sqldb-logger"
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

	loggerAdapter := slogadapter.New(logger)
	pool = sqldblogger.OpenDriver(dbDSN, pool.Driver(), loggerAdapter)

	return pool, nil
}
