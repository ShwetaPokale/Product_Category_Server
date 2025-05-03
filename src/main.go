package main

import (
	"log"
	"net/http"

	"github.com/lenovo/Product_Category_Server/src/config"
	"github.com/lenovo/Product_Category_Server/src/controllers/auth"
	authmiddleware "github.com/lenovo/Product_Category_Server/src/middlewares/auth"
	"github.com/lenovo/Product_Category_Server/src/models/user"
	authroutes "github.com/lenovo/Product_Category_Server/src/routes/auth"
)

func main() {
	// Load database configuration
	dbConfig := config.NewDBConfig()

	// Initialize database connection
	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := user.NewUserRepository(db)

	// Initialize middleware
	authMiddleware := authmiddleware.NewAuthMiddleware()

	// Initialize controllers
	authController := auth.NewAuthController(userRepo, authMiddleware)

	// Initialize router
	mux := http.NewServeMux()

	// Setup routes
	authroutes.SetupAuthRoutes(mux, authController, authMiddleware)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
