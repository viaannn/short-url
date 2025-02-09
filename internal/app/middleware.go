package app

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the request method, URL path and the time it took to process the request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		next.ServeHTTP(writer, request)
		log.Printf("%s %s %v", request.Method, request.URL.Path, time.Since(start))
	})
}
