package main

import (
	"github.com/AndikaSaputra27/booking-system/internal/config"
	"github.com/AndikaSaputra27/booking-system/internal/controller"
	"github.com/AndikaSaputra27/booking-system/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Public routes
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controller.GetUserProfile)
	}

	config.ConnectDB() // pastikan config.DB sudah terhubung

	r.Run(":8080")
}
