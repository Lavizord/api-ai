package auth

import (
	"api-ai/middleware"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Method to generate token, uses secret stored in middleware.
//
// TODO: Review the claims.
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"aud": "audience",
		"iss": "https://issuer/",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middleware.JwtSecret)
}
