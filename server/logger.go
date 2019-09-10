package server

import (
	"log"
	"net/http"
)

type logger struct {
	logger *log.Logger
	server *Server
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
	l.server.ServeHTTP(lw, r)
	l.logger.Printf("status:%d method:%s path:%s\n", lw.statusCode, r.Method, r.URL.Path)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *loggingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
