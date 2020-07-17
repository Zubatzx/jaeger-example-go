package http

import (
	"net/http"

	"book/pkg/grace"
)

// BookHandler ...
type BookHandler interface {
	BookHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server *http.Server
	Book   BookHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	return grace.Serve(port, s.Handler())
}
