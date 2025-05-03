package auth

import (
	"net/http"

	auth "github.com/lenovo/Product_Category_Server/src/controllers/auth"
	authmiddleware "github.com/lenovo/Product_Category_Server/src/middlewares/auth"
)

// SetupAuthRoutes configures the authentication routes
func SetupAuthRoutes(mux *http.ServeMux, authController *auth.AuthController, authMiddleware *authmiddleware.AuthMiddleware) {
	// Public routes
	mux.HandleFunc("POST /api/register", authController.Register)
	mux.HandleFunc("POST /api/login", authController.Login)

	// Protected routes
	mux.Handle("POST /api/logout", authMiddleware.Authenticate(http.HandlerFunc(authController.Logout)))
}
