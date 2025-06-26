package middleware

import (
	"net/http"
	"time"

	"api-ai/internal/logger"
)

// LoggingMiddleware logs incoming requests and their duration.
// It wraps the next handler, printing method and path before and after execution.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log request start: method and URL path
		logger.Default.Infof("→ %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log request completion with elapsed time
		logger.Default.Infof("← done in %v", time.Since(start))
	})
}
