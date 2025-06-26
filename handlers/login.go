package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret") // TODO: must match middleware

func Login(w http.ResponseWriter, r *http.Request) {
	// Normally you'd validate username/password here
	claims := jwt.MapClaims{
		"sub": "user_id_123",
		"aud": "audience",
		"iss": "https://issuer/",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Could not sign token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": signed,
	})
}
