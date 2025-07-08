package models

import "net/http"

// LoginRequest represents login input
type LoginRequest struct {
	Username string `json:"username" example:"user1"`
	Password string `json:"password" example:"pass123"`
}

// TokenResponse represents the token output
type TokenResponse struct {
	Token string `json:"token" example:"your.jwt.token"`
}

// ErrorResponse represents an error message
type ErrorResponse struct {
	Error string `json:"error" example:"invalid credentials"`
}

// ErrorHandlingMiddleware: formats known errors consistently
type ErrorResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (e *ErrorResponseWriter) WriteHeader(status int) {
	e.Status = status
	e.ResponseWriter.WriteHeader(status)
}
