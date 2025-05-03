package product

import (
	"net/http"

	product "Product_Category_Server/src/controllers/product"
	authmiddleware "Product_Category_Server/src/middlewares/auth"
)

// SetupProductRoutes configures the product routes
func SetupProductRoutes(mux *http.ServeMux, productController *product.ProductController, authMiddleware *authmiddleware.AuthMiddleware) {
	// Public routes
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			productController.GetAllProducts(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/products/category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			productController.GetProductsByCategory(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Protected routes

	mux.Handle("/api/products/", authMiddleware.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.GetProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))
}
