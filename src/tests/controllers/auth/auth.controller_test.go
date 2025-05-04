package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Product_Category_Server/src/controllers/auth"
	authmiddleware "Product_Category_Server/src/middlewares/auth"
	"Product_Category_Server/src/models/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepositoryInterface
type MockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() user.UserRepositoryInterface {
	return &MockUserRepository{}
}

func (m *MockUserRepository) CreateUser(user *user.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*user.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *user.User) error {
	// Empty implementation since we don't need this for our tests
	return nil
}

func (m *MockUserRepository) DeleteUser(id string) error {
	// Empty implementation since we don't need this for our tests
	return nil
}

func (m *MockUserRepository) ValidateCredentials(username, password string) (*user.User, error) {
	args := m.Called(username, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func TestAuthController_Register(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]string
		mockError      error
		expectedStatus int
	}{
		{
			name: "Success - Register new user",
			requestBody: map[string]string{
				"username": "newuser",
				"password": "password123",
				"email":    "newuser@example.com",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Error - Missing required fields",
			requestBody: map[string]string{
				"username": "newuser",
				// Missing password and email
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Error - Username already exists",
			requestBody: map[string]string{
				"username": "existinguser",
				"password": "password123",
				"email":    "existinguser@example.com",
			},
			mockError:      assert.AnError,
			expectedStatus: http.StatusConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := NewMockUserRepository()
			authMiddleware := authmiddleware.NewAuthMiddleware()
			controller := auth.NewAuthController(mockRepo, authMiddleware)

			// Prepare request body
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			// Mock expectations
			if tt.requestBody["username"] != "" {
				mockRepo.(*MockUserRepository).On("GetUserByUsername", tt.requestBody["username"]).Return(nil, nil)
				if tt.mockError == nil {
					mockRepo.(*MockUserRepository).On("CreateUser", mock.Anything).Return(nil)
				}
			}

			// Execute
			controller.Register(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedStatus == http.StatusCreated {
				var response map[string]interface{}
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Contains(t, response, "token")
				assert.Contains(t, response, "user")
			}
			mockRepo.(*MockUserRepository).AssertExpectations(t)
		})
	}
}

func TestAuthController_Login(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]string
		mockUser       *user.User
		mockError      error
		expectedStatus int
	}{
		{
			name: "Success - Valid credentials",
			requestBody: map[string]string{
				"username": "testuser",
				"password": "correctpassword",
			},
			mockUser: &user.User{
				ID:       "1",
				Username: "testuser",
				Email:    "test@example.com",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:        "Error - Invalid request body",
			requestBody: map[string]string{
				// Missing username and password
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Error - Invalid credentials",
			requestBody: map[string]string{
				"username": "testuser",
				"password": "wrongpassword",
			},
			mockError:      assert.AnError,
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := NewMockUserRepository()
			authMiddleware := authmiddleware.NewAuthMiddleware()
			controller := auth.NewAuthController(mockRepo, authMiddleware)

			// Prepare request body
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			// Mock expectations
			if tt.requestBody["username"] != "" {
				mockRepo.(*MockUserRepository).On("ValidateCredentials", tt.requestBody["username"], tt.requestBody["password"]).Return(tt.mockUser, tt.mockError)
			}

			// Execute
			controller.Login(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Contains(t, response, "token")
				assert.Contains(t, response, "user")
			}
			mockRepo.(*MockUserRepository).AssertExpectations(t)
		})
	}
}

func TestAuthController_Logout(t *testing.T) {
	// Setup
	mockRepo := NewMockUserRepository()
	authMiddleware := authmiddleware.NewAuthMiddleware()
	controller := auth.NewAuthController(mockRepo, authMiddleware)

	req := httptest.NewRequest("POST", "/api/logout", nil)
	w := httptest.NewRecorder()

	// Execute
	controller.Logout(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Logout successful", response["message"])
}
