-- name: GetUser :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username) VALUES (?) RETURNING *;