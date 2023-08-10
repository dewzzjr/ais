package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(logger{w, r}, r)
	})
}

type logger struct {
	w http.ResponseWriter
	r *http.Request
}

func (l logger) WriteHeader(statusCode int) {
	log.Printf("Request: %s %s\tResponse: %d\n", l.r.Method, l.r.RequestURI, statusCode)
	l.w.WriteHeader(statusCode)
}
func (l logger) Header() http.Header {
	return l.w.Header()
}
func (l logger) Write(b []byte) (int, error) {
	return l.w.Write(b)
}
