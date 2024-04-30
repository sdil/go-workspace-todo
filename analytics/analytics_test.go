package main

import (
	"testing"
	"api/modules/todo"
)

var counttests = []struct{
	title string
	in []todo.Todo
	out int
}{
	{title: "1 false", in: []todo.Todo{{ID: 1, Name: "test", Completed: false}}, out: 0},
	{title: "1 true", in: []todo.Todo{{ID: 1, Name: "test", Completed: true}}, out: 1},
	{title: "true and false", in: []todo.Todo{{ID: 1, Name: "test", Completed: false}, {ID: 1, Name: "test", Completed: true}}, out: 1},
}

func TestCountCompleted(t *testing.T) {
	for _, todo := range counttests {
		t.Run(todo.title, func(t *testing.T) {
			if got, expected := countCompleted(todo.in), todo.out; got != expected {
				t.Fatalf("count (%v) returned %q", todo.in, todo.out)
			}
		})
	}
}
