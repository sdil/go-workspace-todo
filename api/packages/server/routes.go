package server

import "api/modules/todo"

func (s *Server) setupRoutes() {
	todo.SetupRoutes(s.server, s.database)
}
