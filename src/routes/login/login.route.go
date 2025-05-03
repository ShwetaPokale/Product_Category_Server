package login

import (
	"net/http"

	"github.com/lenovo/Product_Category_Server/src/controllers/login"
	"github.com/lenovo/Product_Category_Server/src/middlewares/auth"
)

// SetupLoginRoutes configures all login-related routes
func SetupLoginRoutes(mux *http.ServeMux, loginController *login.LoginController, authMiddleware *auth.AuthMiddleware) {
	// Public routes
	mux.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			loginController.Login(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Protected routes
	mux.Handle("/api/logout", authMiddleware.Authenticate(http.HandlerFunc(loginController.Logout)))
}
