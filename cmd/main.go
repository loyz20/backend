package main

import (
	"log"
	"os"

	"backend/config"
	"backend/internal/product"
	"backend/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	// Connect to database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models (optional dev only)
	db.AutoMigrate(
		&product.Product{},
		&user.User{},
	)

	// Initialize Gin router
	router := gin.Default()

	router.Use(CorsMiddleware())
	router.Use(Logger())

	// Global Middleware
	// router.Use(AuthMiddleware())

	// Grouped API routes
	api := router.Group("/api/v1")
	{
		product.RegisterRoutes(api, db)
		user.RegisterRoutes(api, db)
	}

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	router.Run(":" + port)
}
