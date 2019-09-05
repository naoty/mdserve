package server

import (
	"net/http"
)

// Server is a web server providing RESTful API for markdown contents.
type Server struct {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
