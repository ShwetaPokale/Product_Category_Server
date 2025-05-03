package user

import (
	"database/sql"
	"errors"
	"time"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *User) error {
	query := `
		INSERT INTO users (username, password, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	now := time.Now()
	return r.db.QueryRow(
		query,
		user.Username,
		user.Password,
		user.Email,
		now,
		now,
	).Scan(&user.ID)
}

// GetUserByUsername retrieves a user by username
func (r *UserRepository) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, username, password, email, created_at, updated_at
		FROM users
		WHERE username = $1`

	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(user *User) error {
	query := `
		UPDATE users
		SET username = $1, password = $2, email = $3, updated_at = $4
		WHERE id = $5`

	_, err := r.db.Exec(
		query,
		user.Username,
		user.Password,
		user.Email,
		time.Now(),
		user.ID,
	)
	return err
}

// DeleteUser deletes a user from the database
func (r *UserRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// ValidateCredentials checks if the provided username and password match
func (r *UserRepository) ValidateCredentials(username, password string) (*User, error) {
	user, err := r.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// TODO: In a real application, use proper password hashing
	// For now, we're doing a direct comparison
	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
