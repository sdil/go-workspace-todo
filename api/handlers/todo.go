package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"api/modules/todo"
	"api/modules/user"
)

func TodoRoutes(s *gin.Engine, db *sql.DB) {

	s.GET("/todo/:username", func(c *gin.Context) {
		username := c.Param("username")
		user := user.GetUser(db, username)
		todos := todo.GetTodosByUser(db, user.ID)
		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	})

	s.GET("/todos-create", func(c *gin.Context) {
		todo := todo.CreateTodo(db, todo.Todo{
			Name:      "test",
			Completed: false,
			User: 1,
		})
		c.JSON(http.StatusOK, gin.H{
			"todo": todo,
		})
	})
}
