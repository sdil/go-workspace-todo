package todo

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(s *gin.Engine, db *sql.DB) {

	s.GET("/todos", func(c *gin.Context) {
		todos := GetTodos(db)
		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	})

	s.GET("/todos-create", func(c *gin.Context) {
		todo := CreateTodo(db, Todo{
			Name:      "test",
			Completed: false,
		})
		c.JSON(http.StatusOK, gin.H{
			"todo": todo,
		})
	})
}
