package product_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"Product_Category_Server/src/controllers/product"
	productmodel "Product_Category_Server/src/models/product"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository implements the ProductRepository interface for testing
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindByID(id uint) (*productmodel.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Product), args.Error(1)
}

func (m *MockProductRepository) FindAll() ([]productmodel.Product, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]productmodel.Product), args.Error(1)
}

func (m *MockProductRepository) FindByCategory(category string) ([]productmodel.Product, error) {
	args := m.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]productmodel.Product), args.Error(1)
}

func TestProductController_GetAllProducts(t *testing.T) {
	tests := []struct {
		name           string
		mockProducts   []productmodel.Product
		mockError      error
		expectedStatus int
	}{
		{
			name: "Success - Get all products",
			mockProducts: []productmodel.Product{
				{ProductID: 1, ProductName: "Product 1", ProductPrice: 10.99},
				{ProductID: 2, ProductName: "Product 2", ProductPrice: 20.99},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Error - Database error",
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(MockProductRepository)
			controller := product.NewProductController(mockRepo)
			req := httptest.NewRequest("GET", "/api/products", nil)
			w := httptest.NewRecorder()

			// Mock expectations
			mockRepo.On("FindAll").Return(tt.mockProducts, tt.mockError)

			// Execute
			controller.GetAllProducts(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedStatus == http.StatusOK {
				var response []productmodel.Product
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tt.mockProducts, response)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestProductController_GetProduct(t *testing.T) {
	tests := []struct {
		name           string
		productID      string
		mockProduct    *productmodel.Product
		mockError      error
		expectedStatus int
	}{
		{
			name:           "Success - Get product by ID",
			productID:      "1",
			mockProduct:    &productmodel.Product{ProductID: 1, ProductName: "Product 1", ProductPrice: 10.99},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Error - Missing product ID",
			productID:      "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Error - Invalid product ID",
			productID:      "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Error - Product not found",
			productID:      "999",
			mockError:      assert.AnError,
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(MockProductRepository)
			controller := product.NewProductController(mockRepo)
			req := httptest.NewRequest("GET", "/api/products/detail?id="+tt.productID, nil)
			w := httptest.NewRecorder()

			// Mock expectations
			if tt.productID != "" && tt.productID != "invalid" {
				id, _ := strconv.ParseUint(tt.productID, 10, 32)
				mockRepo.On("FindByID", uint(id)).Return(tt.mockProduct, tt.mockError)
			}

			// Execute
			controller.GetProduct(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedStatus == http.StatusOK {
				var response productmodel.Product
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tt.mockProduct, &response)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestProductController_GetProductsByCategory(t *testing.T) {
	tests := []struct {
		name           string
		category       string
		mockProducts   []productmodel.Product
		mockError      error
		expectedStatus int
	}{
		{
			name:     "Success - Get products by category",
			category: "electronics",
			mockProducts: []productmodel.Product{
				{ProductID: 1, ProductName: "Laptop", ProductCategory: "electronics", ProductPrice: 999.99},
				{ProductID: 2, ProductName: "Smartphone", ProductCategory: "electronics", ProductPrice: 699.99},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Error - Missing category",
			category:       "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Error - Database error",
			category:       "electronics",
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(MockProductRepository)
			controller := product.NewProductController(mockRepo)
			req := httptest.NewRequest("GET", "/api/products/category?category="+tt.category, nil)
			w := httptest.NewRecorder()

			// Mock expectations
			if tt.category != "" {
				mockRepo.On("FindByCategory", tt.category).Return(tt.mockProducts, tt.mockError)
			}

			// Execute
			controller.GetProductsByCategory(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedStatus == http.StatusOK {
				var response []productmodel.Product
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tt.mockProducts, response)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
