package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AFelipeTrujillo/homysync-go-api/internal/adapters/handlers"
	"github.com/AFelipeTrujillo/homysync-go-api/internal/adapters/repository"
	"github.com/AFelipeTrujillo/homysync-go-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Starting HomySync Go API...")

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	repository.InitDB()

	// Initialize repositories
	userRepo := repository.NewGormUserRepository(repository.DB)

	// Initialize services
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)

}
