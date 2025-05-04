package product

import (
	"log"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(id uint) (*Product, error)
	FindAll() ([]Product, error)
	FindByCategory(category string) ([]Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindByID(id uint) (*Product, error) {
	var product Product
	err := r.db.Where("product_id = ?", id).First(&product).Error
	if err != nil {
		log.Printf("Error finding product by ID %d: %v", id, err)
	} else {
		log.Printf("Found product: %+v", product)
	}
	return &product, err
}

func (r *productRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		log.Printf("Error finding all products: %v", err)
	} else {
		log.Printf("Found %d products", len(products))
		for i, p := range products {
			log.Printf("Product %d: %+v", i+1, p)
		}
	}
	return products, err
}

func (r *productRepository) FindByCategory(category string) ([]Product, error) {
	var products []Product
	err := r.db.Where("product_category = ?", category).Find(&products).Error
	if err != nil {
		log.Printf("Error finding products by category %s: %v", category, err)
	} else {
		log.Printf("Found %d products in category %s", len(products), category)
		for i, p := range products {
			log.Printf("Product %d: %+v", i+1, p)
		}
	}
	return products, err
}
