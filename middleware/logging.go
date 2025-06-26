package middleware

import (
	"net/http"
	"time"

	"api-ai/internal/logger"

	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// LoggingMiddleware logs incoming requests and their duration.
//
// It wraps the next handler, logs method and path before and after execution.
// Also logs the user and ip
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ip := r.RemoteAddr
		user := "unknown"
		if claims, ok := r.Context().Value("user").(*validator.ValidatedClaims); ok {
			user = claims.RegisteredClaims.Subject
		}

		logger.Default.Infof("→ %s %s | user=%s | ip=%s", r.Method, r.URL.Path, user, ip)
		next.ServeHTTP(w, r)
		logger.Default.Infof("← %s %s | done in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
