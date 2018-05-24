package server

import (
	"net/http"
)

// Server provides http service for the labchat service.
type Server struct {
	// TODO: implementation.
}

// NewServer creates a new labchat server with the given configuration.
func NewServer(cfg *Config) (srv *Server, err error) {
	// TODO: implementation.
	return nil, nil
}

// Start runs the server and starts handling for incoming requests. All the
// long-running server functionality should be implemented in goroutines.
func (s *Server) Start() {
	// TODO: implementation.
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":80", nil)
}
