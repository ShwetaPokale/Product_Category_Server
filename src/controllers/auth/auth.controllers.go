package auth

import (
	"encoding/json"
	"net/http"

	authmiddleware "Product_Category_Server/src/middlewares/auth"
	"Product_Category_Server/src/models/user"
)

type AuthController struct {
	userRepo user.UserRepositoryInterface
	auth     *authmiddleware.AuthMiddleware
}

func NewAuthController(userRepo user.UserRepositoryInterface, auth *authmiddleware.AuthMiddleware) *AuthController {
	return &AuthController{
		userRepo: userRepo,
		auth:     auth,
	}
}

// Register handles user registration requests
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if newUser.Username == "" || newUser.Password == "" || newUser.Email == "" {
		http.Error(w, "Username, password, and email are required", http.StatusBadRequest)
		return
	}

	// Check if username already exists
	existingUser, _ := c.userRepo.GetUserByUsername(newUser.Username)
	if existingUser != nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Create the user
	if err := c.userRepo.CreateUser(&newUser); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Generate token
	token := c.auth.GenerateMockToken()

	// Return success response with token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Registration successful",
		"token":   token,
		"user": map[string]interface{}{
			"id":       newUser.ID,
			"username": newUser.Username,
			"email":    newUser.Email,
		},
	})
}

// Login handles user login requests
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate credentials against the database
	user, err := c.userRepo.ValidateCredentials(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate token
	token := c.auth.GenerateMockToken()

	// Return success response with token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"token":   token,
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Logout handles user logout requests
func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: In a real application, invalidate the token
	// For now, just return success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}
