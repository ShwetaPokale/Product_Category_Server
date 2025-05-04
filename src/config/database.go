package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Product_Category_Server/src/models/product"
	"Product_Category_Server/src/models/user"
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
		Password: getEnv("DB_PASSWORD", "root"),
		DBName:   getEnv("DB_NAME", "my_database"),
	}
}

func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}

func InitDB(config *DBConfig) (*gorm.DB, error) {
	dsn := config.GetDSN()
	log.Printf("Connecting to database with DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	log.Println("Successfully connected to database")

	// Auto-migrate the schema
	err = db.AutoMigrate(&user.User{}, &product.Product{})
	if err != nil {
		return nil, fmt.Errorf("error migrating database: %v", err)
	}

	log.Println("Database migration completed")

	// Test the connection by counting products
	var count int64
	if err := db.Model(&product.Product{}).Count(&count).Error; err != nil {
		log.Printf("Error counting products: %v", err)
	} else {
		log.Printf("Current number of products in database: %d", count)
	}

	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
