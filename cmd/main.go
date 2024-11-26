package main

import (
	"fmt"
	"log"
	"muyi2905/civic/backend/models"
	"muyi2905/civic/backend/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDb initializes the database connection
func InitDb() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading environment variables: %v", err)
	} else {
		log.Println("Environment variables loaded successfully")
	}

	// Get the DSN (Data Source Name) from environment variables
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("Database DSN not found in environment variables.")
	}

	log.Printf("Connecting to database: %s", dsn)

	// Open the database connection using gorm
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	fmt.Println("Database connection successful")
}

func main() {
	// Initialize the database connection
	InitDb()

	// Migrate the database schema for the ElectedOfficial model
	if err := DB.AutoMigrate(&models.ElectedOfficial{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	// Set up the Gin router
	r := gin.Default()

	// Set up the routes
	routes.UserRoutes(DB, r)     // User routes (you might already have these)
	routes.(DB, r) // Elected Official routes

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
