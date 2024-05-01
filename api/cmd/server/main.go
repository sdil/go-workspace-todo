package main

import (
	"api/modules/todo"
	"api/modules/user"
	"api/packages/server"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	log.Println("Starting server")

	db, err := sql.Open("sqlite3", "../data.sqlite3")
	if err != nil {
		log.Fatalln("Unable to start sqlite")
	}
	todo.CreateTables(db)
	user.CreateTables(db)

	s := server.NewServer(
		server.NewServerOptions{
			Database: db,
		},
	)

	s.Start()
}
