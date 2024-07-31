-- name: CreateTask :exec
INSERT INTO tasks (name, status, due_date, priority) VALUES (?, ?, ?, ?);

-- name: GetTasks :many
SELECT id, name, status, due_date, priority FROM tasks WHERE status = ?;

