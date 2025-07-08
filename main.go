package main

import (
	"api-ai/db"
	_ "api-ai/docs"
	"api-ai/ent"
	"api-ai/internal/logger"
	"api-ai/internal/services"
	"api-ai/middleware"
	"api-ai/routes"
	"api-ai/secrets"
	"log"
	"net/http"
)

var dbcli *ent.Client

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger.Default.Info("Api initializing...")
	dbcli = db.NewClient()
	logger.Default.Info("Entgo db initialized...")

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
	logger.Default.Info("JWTMiddleware initialized...")

	fh, err := services.NewFileHandler(dbcli) // Inits the file handler that will ne used in the endpoints.
	if err != nil {
		logger.Default.Fatalf(err.Error())
	}
	logger.Default.Info("FileHandler initialized...")

	r := routes.RegisterRoutes(fh) // Registers routs with services provided.

	logger.Default.Info("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
