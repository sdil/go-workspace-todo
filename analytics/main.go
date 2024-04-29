package main

import (
	"api/modules/todo"
	"log"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../data.sqlite3")
	if err != nil {
		log.Fatalln("Unable to start sqlite")
	}
	todos := todo.GetTodos(db)
	completed := countCompleted(todos)
	log.Println(completed)
}
