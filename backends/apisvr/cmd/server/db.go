package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		return nil, fmt.Errorf("DB_DSN is not set")
	}

	pool, err := sql.Open("mysql", dbDSN)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
