package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"2FA/internal/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.Status(200)
	})

	r.POST("/login", handlers.LoginHandler)
	r.POST("/verify", handlers.VerifyLoginHandler)
	r.POST("/reset_password", handlers.ResetPasswordHandler)
	r.POST("/verify_reset_password", handlers.VerifyResetHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8084"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
