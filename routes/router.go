package routes

import (
	"api-ai/handlers"
	"api-ai/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoveryMiddleware)
	r.Use(middleware.CORSMiddleware())

	// Public routes
	// /auth routes
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	// /api protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.HandleFunc("/upload", handlers.UploadPDF).Methods("POST")

	return r
}
