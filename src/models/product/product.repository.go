package product

import "gorm.io/gorm"

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
	return &product, err
}

func (r *productRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByCategory(category string) ([]Product, error) {
	var products []Product
	err := r.db.Where("category = ?", category).Find(&products).Error
	return products, err
}
