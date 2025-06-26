package handlers

import (
	"api-ai/middleware"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Login method that supplies JWT token with duration and claims to be used on the api.
//
// TODO: validate username/password here.
// TODO: Review claims.
func Login(w http.ResponseWriter, r *http.Request) {
	claims := jwt.MapClaims{
		"sub": "user_id_123",
		"aud": "audience",
		"iss": "https://issuer/",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(middleware.JwtSecret)
	if err != nil {
		http.Error(w, "Could not sign token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": signed,
	})
}
