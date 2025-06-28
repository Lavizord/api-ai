package middleware

import (
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
)

// SwaggerHandler returns a protected Swagger UI handler
func SwaggerHandler() http.Handler {
	handler := httpSwagger.WrapHandler

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow access only in dev/test env
		env := os.Getenv("ENV")
		if env == "PROD" {
			http.NotFound(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
