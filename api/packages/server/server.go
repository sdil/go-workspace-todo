package server

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type NewServerOptions struct {
	Database *sql.DB
}

type Server struct {
	database *sql.DB
	server   *gin.Engine
}

func NewServer(opts NewServerOptions) *Server {
	log.Println("Starting server")
	server := gin.Default()

	log.Println("Setting up routes")

	return &Server{
		database: opts.Database,
		server:   server,
	}
}

func (s *Server) Start() {
	s.setupRoutes()
	s.server.Run("localhost:8080")
}
