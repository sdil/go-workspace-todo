package server

import "api/packages/server/routes"

func (s *Server) setupRoutes() {
	routes.TodoRoutes(s.server, s.database)
}
