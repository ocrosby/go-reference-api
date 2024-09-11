package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request details
		log.Printf("Received request: method=%s url=%s time=%s", r.Method, r.URL.String(), start.Format(time.RFC3339))

		// Create a response recorder to capture the status code
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(rr, r)

		// Log the response time and status code
		log.Printf("Completed request: method=%s url=%s responseTime=%s statusCode=%d", r.Method, r.URL.String(), time.Since(start).String(), rr.statusCode)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}
