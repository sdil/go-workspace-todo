package todo

import (
	"context"
	"database/sql"
	_ "embed"
)

func GetTodos(db *sql.DB) []Todo {
	ctx := context.Background()
	queries := New(db)
	todos, _ := queries.ListTodo(ctx)
	return todos
}

func CreateTodo(db *sql.DB, t Todo) Todo {
	ctx := context.Background()
	queries := New(db)
	todo, _ := queries.CreateTodo(ctx, t.Name)
	return todo
}

//go:embed schema.sql
var ddl string

func CreateTables(db *sql.DB) {
	ctx := context.Background()
	db.ExecContext(ctx, ddl)
}
