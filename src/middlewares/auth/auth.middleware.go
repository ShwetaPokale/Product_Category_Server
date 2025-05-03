package auth

import (
	"net/http"
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
		// TODO: Implement authentication logic
		// For example, check for valid JWT token
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication passes, call the next handler
		next.ServeHTTP(w, r)
	})
}
