package middleware

import (
	"encoding/json"
	"net/http"

	"api-ai/internal/logger"
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

// ErrorHandlingMiddleware: formats known errors consistently
type errorResponseWriter struct {
	http.ResponseWriter
	status int
}

func (e *errorResponseWriter) WriteHeader(status int) {
	e.status = status
	e.ResponseWriter.WriteHeader(status)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		erw := &errorResponseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(erw, r)

		// Customize JSON error formatting
		if erw.status >= 400 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ErrorResponse{
				Error: http.StatusText(erw.status),
			})
		}
	})
}
