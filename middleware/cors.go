package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSMiddleware returns a middleware function that applies CORS settings to HTTP handlers.
//
// Returned object should be used in the router.
func CORSMiddleware() func(http.Handler) http.Handler {
	return handlers.CORS(
		// Allow all origins for now; restrict this in production for security.
		handlers.AllowedOrigins([]string{"*"}),

		// Allowed HTTP methods on the resource.
		//handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),

		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowCredentials(),
	)
}
