package routes

import (
	"api-ai/handlers"
	"api-ai/internal/services"
	"api-ai/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(fh *services.FileHandler) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoveryMiddleware)
	r.Use(middleware.ErrorHandlingMiddleware)
	r.Use(middleware.CORSMiddleware())

	r.PathPrefix("/swagger/").Handler(middleware.SwaggerHandler())

	// Public routes
	// /auth routes
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET", "OPTIONS")
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", handlers.Login).Methods("POST", "OPTIONS")

	// Protected routes
	// /api protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		handlers.UploadPDF(w, r, fh)
	}).Methods("POST", "OPTIONS")

	return r
}
