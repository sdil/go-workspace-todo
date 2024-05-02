package handlers

import (
	"api/modules/user"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(s *gin.Engine, db *sql.DB) {

	s.GET("/user/:username", func(c *gin.Context) {
		username := c.Param("username")
		user := user.GetUser(db, username)
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	s.GET("/create-user/:username", func(c *gin.Context) {
		username := c.Param("username")
		user := user.CreateUser(db, user.User{
			Username: username,
		})
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})
}
