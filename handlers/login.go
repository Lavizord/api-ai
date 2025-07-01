package handlers

import (
	"api-ai/middleware"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Login godoc
// @Summary      Authenticate user and return JWT token
// @Description  Takes a username and password and returns a JWT token if valid
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest     true  "User credentials"
// @Success      200          {object}  TokenResponse
// @Failure      500          {object}  ErrorResponse
// @Router       /auth/login [post]
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
