package login

import (
	"encoding/json"
	"net/http"

	"github.com/lenovo/Product_Category_Server/src/models/user"
)

// LoginController handles login-related requests
type LoginController struct {
	userRepo *user.UserRepository
}

// NewLoginController creates a new instance of LoginController
func NewLoginController(userRepo *user.UserRepository) *LoginController {
	return &LoginController{
		userRepo: userRepo,
	}
}

// Login handles user login requests
func (c *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Implement actual authentication logic
	// For now, just return a mock response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   "mock-jwt-token",
	})
}

// Logout handles user logout requests
func (c *LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logout logic
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}
