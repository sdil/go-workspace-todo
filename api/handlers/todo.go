package handlers

import (
	"api/modules/todo"
	"api/modules/user"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(s *gin.Engine, db *sql.DB) {

	s.GET("/todo/:username", func(c *gin.Context) {
		tx, err := db.Begin()
		log.Println("start txn")
		if err != nil {
			log.Fatalln("Failed to create transaction")
		}
		defer tx.Rollback()

		username := c.Param("username")
		user := user.GetUser(db, username)
		todos := todo.GetTodosByUser(db, user.ID)

		tx.Commit()
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
