package main

import (
	"api/modules/todo"
)

func countCompleted(t []todo.Todo) int {
	completed := 0
	for _, todo := range t {
		if todo.Completed {
			completed += 1
		}
	}

	return completed
}
