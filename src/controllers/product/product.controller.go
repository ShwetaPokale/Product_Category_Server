package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"Product_Category_Server/src/models/product"
)

type ProductController struct {
	productRepo product.ProductRepository
}

func NewProductController(productRepo product.ProductRepository) *ProductController {
	return &ProductController{
		productRepo: productRepo,
	}
}

func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.productRepo.FindAll()
	if err != nil {
		log.Printf("Error getting all products: %v", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("Error encoding products: %v", err)
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
	}
}

func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL
	path := r.URL.Path
	idStr := path[strings.LastIndex(path, "/")+1:]

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Invalid product ID: %v", err)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := c.productRepo.FindByID(uint(id))
	if err != nil {
		log.Printf("Error finding product: %v", err)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Printf("Error encoding product: %v", err)
		http.Error(w, "Failed to encode product", http.StatusInternalServerError)
	}
}

func (c *ProductController) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {

	category := r.URL.Query().Get("name")
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	products, err := c.productRepo.FindByCategory(category)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
