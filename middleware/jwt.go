package middleware

import (
	"context"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

var JWTMiddlewareInstance *jwtmiddleware.JWTMiddleware
var JwtSecret []byte

func InitJWTMiddleware(secret []byte, issuer string, audience []string) error {
	JwtSecret = secret

	keyFunc := func(ctx context.Context) (interface{}, error) {
		return secret, nil
	}

	jwtValidator, err := validator.New(keyFunc, validator.HS256, issuer, audience)
	if err != nil {
		return err
	}

	JWTMiddlewareInstance = jwtmiddleware.New(jwtValidator.ValidateToken)
	return nil
}

func JWTMiddleware(next http.Handler) http.Handler {
	return JWTMiddlewareInstance.CheckJWT(next)
}
