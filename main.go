package main

import (
	_ "api-ai/docs"
	"api-ai/internal/logger"
	"api-ai/middleware"
	"api-ai/routes"
	"api-ai/secrets"
	"log"
	"net/http"
)

func main() {
	logger.Default.Info("Api initializing...")

	// Load JWT secret from AWS Secrets Manager
	jwtSecretStr, err := secrets.GetSecretValue("your-secret-id")
	if err != nil {
		// TODO: This cant stay like this for long.
		logger.Default.Errorf("Failed to get aws secret, changing to dev default, aws message: %v", err)
		jwtSecretStr = "dev-secret-placeholder"
	}
	err = middleware.InitJWTMiddleware([]byte(jwtSecretStr), "https://issuer/", []string{"audience"})
	if err != nil {
		logger.Default.Fatalf("Issue initializing JWT Middleware: %v", err)
	}
	r := routes.RegisterRoutes()
	logger.Default.Info("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
