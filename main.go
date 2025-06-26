package main

import (
	"api-ai/internal/logger"
	"api-ai/middleware"
	"api-ai/routes" // import your router package
	"log"
	"net/http"
)

func main() {
	logger.Default.Info("Api initializing...")
	// TODO: These need to be the same used in the other packages, aws secrets or a config?
	err := middleware.InitJWTMiddleware([]byte("your-secret"), "https://issuer/", []string{"audience"})
	if err != nil {
		log.Fatal(err)
	}
	r := routes.RegisterRoutes() // get configured router
	logger.Default.Info("Registered routes...")
	logger.Default.Info("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
