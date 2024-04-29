package main

import (
	"api/modules/todo"
	"log"

	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../data.sqlite3")
	if err != nil {
		log.Fatalln("Unable to start sqlite")
	}
	log.Println(todo.GetTodos(db))
}