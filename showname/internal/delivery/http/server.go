package http

import (
	"net/http"

	"showname/pkg/grace"
)

// ShownameHandler ...
type ShownameHandler interface {
	ShownameHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server   *http.Server
	Showname ShownameHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	return grace.Serve(port, s.Handler())
}
