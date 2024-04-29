package routes

import (
	"api/modules/todo"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TodoRoutes(s *gin.Engine, db *sql.DB) {

	s.GET("/todos", func(c *gin.Context) {
		todos := todo.GetTodos(db)
		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	})

	s.GET("/todos-create", func(c *gin.Context) {
		todo := todo.CreateTodo(db, todo.Todo{
			Name:      "test",
			Completed: false,
		})
		c.JSON(http.StatusOK, gin.H{
			"todo": todo,
		})
	})
}
