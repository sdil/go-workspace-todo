-- name: GetTodo :one
SELECT * FROM todos WHERE id = ? LIMIT 1;

-- name: ListTodo :many
SELECT * FROM todos;

-- name: ListTodoByUser :many
SELECT * FROM todos WHERE user = ?;

-- name: CreateTodo :one
INSERT INTO todos (name, completed) VALUES (?, 0) RETURNING *;