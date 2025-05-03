package auth

import (
	"net/http"

	auth "Product_Category_Server/src/controllers/auth"
	authmiddleware "Product_Category_Server/src/middlewares/auth"
)

// SetupAuthRoutes configures the authentication routes
func SetupAuthRoutes(mux *http.ServeMux, authController *auth.AuthController, authMiddleware *authmiddleware.AuthMiddleware) {
	// Public routes
	mux.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			authController.Register(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			authController.Login(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Protected routes
	mux.Handle("/api/logout", authMiddleware.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			authController.Logout(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))
}
