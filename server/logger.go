package server

import (
	"log"
	"net/http"
	"time"
)

type logger struct {
	logger *log.Logger
	server *Server
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	lw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
	l.server.ServeHTTP(lw, r)

	timestamp := now.Format("2006-01-02T15:04:05.000Z07:00")
	l.logger.Printf(
		"%s %s %s %d %dμs\n",
		timestamp,
		r.Method,
		r.URL.Path,
		lw.statusCode,
		time.Since(now).Microseconds(),
	)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *loggingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
