package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // Password is not exposed in JSON
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*User, error) {
	query := `
		SELECT id, username, password, email, created_at, updated_at
		FROM users
		WHERE username = $1
	`
	user := &User{}
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Create(user *User) error {
	query := `
		INSERT INTO users (username, password, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	return r.db.QueryRow(
		query,
		user.Username,
		user.Password,
		user.Email,
		time.Now(),
		time.Now(),
	).Scan(&user.ID)
}

func (r *UserRepository) Update(user *User) error {
	query := `
		UPDATE users
		SET username = $1, password = $2, email = $3, updated_at = $4
		WHERE id = $5
	`
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

func (r *UserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
