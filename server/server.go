package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/naoty/mdserve/contents"
)

// Server is a web server providing RESTful API for markdown contents.
type Server struct {
	routes map[string]http.Handler
}

// New returns a new Server.
func New(path string) *Server {
	path = normalizedContentsPath(path)

	routes := map[string]http.Handler{}
	routes[path] = contentsHandler()
	routes[path+"index.json"] = contentsHandler()

	for _, content := range contents.Index() {
		path := content["filepath"].(string)
		normalizedPath := normalizeContentPath(path)
		routes[normalizedPath] = contentHandler(path)
	}

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
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)

		c := contents.Index()
		encoder.Encode(c)
	})
}

func contentHandler(path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)

		c, ok := contents.Get(path)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		encoder.Encode(c)
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

func normalizedContentsPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}

	if !strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s/", path)
	}

	return path
}

func normalizeContentPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}

	return strings.Replace(path, ".md", ".json", 1)
}
