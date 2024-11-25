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

func InitDb() {

	if err := godotenv.Load(); err != nil {
		log.Printf("error loading environment variable : %v", err)
	} else {
		log.Println("Environment variables loaded successfully")
	}
	dsn := os.Getenv("DSN")
	log.Printf("DSN: %s", dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	fmt.Println("database connection successful")
}

func main() {
	InitDb()
	r := gin.Default()
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to auto migrate database :%v", err)
	}
	routes.UserRoutes(DB, r)
	r.Run(":8080")
}
