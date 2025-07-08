package middleware

import (
	"encoding/json"
	"net/http"

	"api-ai/internal/logger"
	"api-ai/internal/models"
)

// RecoveryMiddleware: only logs panics
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Default.Errorf("panic recovered: %v", rec)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		erw := &models.ErrorResponseWriter{ResponseWriter: w, Status: http.StatusOK}
		next.ServeHTTP(erw, r)

		// Customize JSON error formatting
		if erw.Status >= 400 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: http.StatusText(erw.Status),
			})
		}
	})
}
