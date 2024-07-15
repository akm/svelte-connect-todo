package migrations

import (
	"context"
	"database/sql"
	"dbmigrations/helpers"
)

func init() {
	helpers.DemoData.AddFileNameMigrationContext(upGo, downGo)
}

func upGo(ctx context.Context, tx *sql.Tx) error {
	statement :=
		"INSERT INTO tasks (id, name, status) VALUES " +
			"(1, 'List items', 'done')," +
			"(2, 'Buy items', 'todo');"
	if _, err := tx.Exec(statement); err != nil {
		return err
	}
	return nil
}

func downGo(ctx context.Context, tx *sql.Tx) error {
	statement := "DELETE FROM tasks WHERE id IN (1,2);"
	if _, err := tx.Exec(statement); err != nil {
		return err
	}

	return nil
}
