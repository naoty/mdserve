package server

import (
	"net/http"
)

// Server is a web server providing RESTful API for markdown contents.
type Server struct {
	routes map[string]http.Handler
}

// New returns a new Server.
func New() *Server {
	routes := map[string]http.Handler{}
	routes["/contents"] = contentsHandler()

	return &Server{
		routes: routes,
	}
}

func contentsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler, ok := s.routes[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler.ServeHTTP(w, r)
}
