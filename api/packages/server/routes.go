package server

import "api/handlers"

func (s *Server) setupRoutes() {
	handlers.TodoRoutes(s.server, s.database)
}
