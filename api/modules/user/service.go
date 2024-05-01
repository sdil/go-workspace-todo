package user

import (
	"context"
	"database/sql"
	_ "embed"
)

func GetUser(db *sql.DB, username string) User {
	ctx := context.Background()
	queries := New(db)
	user, _ := queries.GetUser(ctx, username)
	return user
}

func CreateUser(db *sql.DB, user User) User {
	ctx := context.Background()
	queries := New(db)
	new_user, _ := queries.CreateUser(ctx, user.Username)
	return new_user
}

//go:embed schema.sql
var ddl string

func CreateTables(db *sql.DB) {
	ctx := context.Background()
	db.ExecContext(ctx, ddl)
}
