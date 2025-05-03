package services

import (
	"golang_backend/app/models"
	"golang_backend/app/repositories"
)

type ProductService interface {
	GetProductByID(id uint) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	GetProductDetails(id uint) (*models.Product, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productService) GetProductsByCategory(category string) ([]models.Product, error) {
	return s.productRepo.FindByCategory(category)
}

func (s *productService) GetProductDetails(id uint) (*models.Product, error) {
	return s.productRepo.FindByID(id)
}
