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

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=0629 dbname=simple_blog port=5432 sslmode=disable"

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