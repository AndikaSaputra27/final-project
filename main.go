package main

import (
	"github.com/AndikaSaputra27/booking-system/internal/controller"
	"github.com/AndikaSaputra27/booking-system/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Membuat instance dari AuthService
	authService := service.NewAuthService()

	// Membuat instance dari AuthController dengan mengirimkan AuthService
	authController := controller.NewAuthController(authService)

	// Setup Gin Router
	router := gin.Default()

	// Endpoint login
	router.POST("/login", authController.Login)

	// Jalankan server
	router.Run(":8080")
}
