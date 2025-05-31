package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func InitDB() {
	var err error
	
	// Get database connection details from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "0629")
	dbName := getEnv("DB_NAME", "simple_blog")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName, dbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")

	// Create tables using raw SQL queries
	err = CreateTables()
	if err != nil {
		log.Fatal("Failed to create tables:", err)
		os.Exit(1)
	}

	// Auto migrate for any additional changes in the model
	err = DB.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		os.Exit(1)
	}

	fmt.Println("Database tables created and migrated successfully!")
} 