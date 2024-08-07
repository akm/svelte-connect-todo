// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tasks.sql

package models

import (
	"context"
	"database/sql"
)

const countTasks = `-- name: CountTasks :one
SELECT count(*) FROM tasks
`

func (q *Queries) CountTasks(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTasks)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTask = `-- name: CreateTask :execresult
INSERT INTO tasks (
  name, status
) VALUES (
  ?, ?
)
`

type CreateTaskParams struct {
	Name   string
	Status TasksStatus
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTask, arg.Name, arg.Status)
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?
`

func (q *Queries) DeleteTask(ctx context.Context, id uint64) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTask = `-- name: GetTask :one
SELECT id, created_at, updated_at, name, status FROM tasks 
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id uint64) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Status,
	)
	return i, err
}

const getTaskForUpdate = `-- name: GetTaskForUpdate :one
SELECT id, created_at, updated_at, name, status FROM tasks 
WHERE id = ? LIMIT 1 FOR UPDATE
`

func (q *Queries) GetTaskForUpdate(ctx context.Context, id uint64) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskForUpdate, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Status,
	)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, created_at, updated_at, name, status FROM tasks
ORDER BY id ASC
`

func (q *Queries) ListTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :exec
UPDATE tasks
SET name = ?, status = ?
WHERE id = ?
`

type UpdateTaskParams struct {
	Name   string
	Status TasksStatus
	ID     uint64
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateTask, arg.Name, arg.Status, arg.ID)
	return err
}
