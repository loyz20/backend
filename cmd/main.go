package main

import (
	"fmt"
	"log"
	"os"

	"backend/config"
	"backend/internal/customer"
	"backend/internal/order"
	"backend/internal/product"
	"backend/internal/sale"
	"backend/internal/shared"
	"backend/internal/supplier"
	"backend/internal/user"
	"backend/migrations"

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

	err = migrations.RunAll(db)
	if err != nil {
		log.Fatal("Migration error: ", err)
	}

	err = migrations.Seed(db)
	if err != nil {
		log.Fatal("Seeding error: ", err)
	}

	fmt.Println("Migration & seeding completed.")

	// Initialize Gin router
	router := gin.Default()

	router.Use(shared.CORSMiddleware())
	router.Use(shared.Logger())

	// Global Middleware
	// router.Use(AuthMiddleware())

	// Grouped API routes
	api := router.Group("/api/v1")
	{
		user.RegisterRoutes(api, db)
		auth := router.Group("/api/v1/auth")
		{
			auth.Use(shared.JWTAuthMiddleware())
			product.RegisterRoutes(auth, db)
			customer.RegisterRoutes(auth, db)
			supplier.RegisterRoutes(auth, db)
			order.RegisterRoutes(auth, db)
			sale.RegisterRoutes(auth, db)

		}
	}

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	router.Run(":" + port)
}
