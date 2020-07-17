package http

import (
	"net/http"

	"showtime/pkg/grace"
)

// ShowtimeHandler ...
type ShowtimeHandler interface {
	ShowtimeHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server   *http.Server
	Showtime ShowtimeHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	return grace.Serve(port, s.Handler())
}
