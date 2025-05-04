package user

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *User) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	return r.db.Create(user).Error
}

// GetUserByUsername retrieves a user by username
func (r *UserRepository) GetUserByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(user *User) error {
	user.UpdatedAt = time.Now()
	return r.db.Save(user).Error
}

// DeleteUser deletes a user from the database
func (r *UserRepository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}

// ValidateCredentials checks if the provided username and password match
func (r *UserRepository) ValidateCredentials(username, password string) (*User, error) {
	user, err := r.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// TODO: Use password hashing in production
	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
