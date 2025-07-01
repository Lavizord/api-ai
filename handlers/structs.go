package handlers

// TODO: This should be something global. The error response should be the same as the error middleware response.

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
