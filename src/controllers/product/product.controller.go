package product

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := c.productRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	products, err := c.productRepo.FindByCategory(category)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}


