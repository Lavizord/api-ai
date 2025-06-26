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
	r.Use(middleware.JWTMiddleware)
	r.HandleFunc("/upload", handlers.UploadPDF).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	return r
}
