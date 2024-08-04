-- name: CountTasks :one
SELECT count(*) FROM tasks;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY id ASC;

-- name: GetTask :one
SELECT * FROM tasks 
WHERE id = ? LIMIT 1;

-- name: GetTaskForUpdate :one
SELECT * FROM tasks 
WHERE id = ? LIMIT 1 FOR UPDATE;

-- name: CreateTask :execresult
INSERT INTO tasks (
  name, status
) VALUES (
  ?, ?
);

-- name: UpdateTask :exec
UPDATE tasks
SET name = ?, status = ?
WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;
