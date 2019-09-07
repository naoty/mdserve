package server

import (
	"encoding/json"
	"log"
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

// WithLogger returns a HTTP handler wrapping Server with logger.
func (s *Server) WithLogger(l *log.Logger) http.Handler {
	return &logger{
		logger: l,
		server: s,
	}
}

func contentsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contents := []map[string]string{}

		contents = append(contents, map[string]string{
			"title": "Test1",
			"body":  "this is test content",
		})
		contents = append(contents, map[string]string{
			"title": "Test2",
			"body":  "this is test content",
		})

		data, err := json.Marshal(contents)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
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
