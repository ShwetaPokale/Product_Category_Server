package auth

import (
	"encoding/json"
	"net/http"
	"strings"
)

// AuthMiddleware handles authentication for protected routes
type AuthMiddleware struct {
	// Add any dependencies here
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// Authenticate is the middleware function that checks for valid authentication
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Check if the header has the correct format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		// Get the token
		token := parts[1]

		// TODO: In a real application, validate the JWT token here
		// For now, we'll just check if it's our mock token
		if token != "mock-jwt-token" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If authentication passes, call the next handler
		next.ServeHTTP(w, r)
	})
}

// GenerateMockToken generates a mock JWT token for testing
func (m *AuthMiddleware) GenerateMockToken() string {
	return "mock-jwt-token"
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// sendErrorResponse sends a JSON error response
func sendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
