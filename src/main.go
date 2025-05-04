package main

import (
	"log"
	"net/http"

	"Product_Category_Server/src/config"
	authmiddleware "Product_Category_Server/src/middlewares/auth"
	corsmiddleware "Product_Category_Server/src/middlewares/cors"

	"Product_Category_Server/src/controllers/auth"
	"Product_Category_Server/src/controllers/product"
	PRODUCT "Product_Category_Server/src/models/product"
	"Product_Category_Server/src/models/user"
	authroutes "Product_Category_Server/src/routes/auth"
	productroutes "Product_Category_Server/src/routes/product"
)

func main() {
	// Load database configuration
	dbConfig := config.NewDBConfig()

	// Initialize database connection
	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repositories
	userRepo := user.NewUserRepository(db)
	productRepo := PRODUCT.NewProductRepository(db)

	// Initialize middleware
	authMiddleware := authmiddleware.NewAuthMiddleware()
	corsMiddleware := corsmiddleware.NewCORSMiddleware()

	// Initialize controllers
	authController := auth.NewAuthController(userRepo, authMiddleware)
	productController := product.NewProductController(productRepo)

	// Initialize router
	mux := http.NewServeMux()

	// Setup routes with CORS middleware
	handler := corsMiddleware.HandleCORS(mux)
	authroutes.SetupAuthRoutes(mux, authController, authMiddleware)
	productroutes.SetupProductRoutes(mux, productController, authMiddleware)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
