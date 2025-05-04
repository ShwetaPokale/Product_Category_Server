package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "post123"),
		DBName:   getEnv("DB_NAME", "my_database"),
	}
}

func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}

func InitDB(config *DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
