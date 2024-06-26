// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package todo

import (
	"context"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (name, completed) VALUES (?, 0) RETURNING id, name, completed, user
`

func (q *Queries) CreateTodo(ctx context.Context, name string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, name)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.User,
	)
	return i, err
}

const getTodo = `-- name: GetTodo :one
SELECT id, name, completed, user FROM todos WHERE id = ? LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.User,
	)
	return i, err
}

const listTodo = `-- name: ListTodo :many
SELECT id, name, completed, user FROM todos
`

func (q *Queries) ListTodo(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Completed,
			&i.User,
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

const listTodoByUser = `-- name: ListTodoByUser :many
SELECT id, name, completed, user FROM todos WHERE user = ?
`

func (q *Queries) ListTodoByUser(ctx context.Context, user interface{}) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodoByUser, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Completed,
			&i.User,
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
