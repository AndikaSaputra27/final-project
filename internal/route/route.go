package route

import (
	"github.com/AndikaSaputra27/booking-system/internal/controller"
	"github.com/AndikaSaputra27/booking-system/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	// Protected routes - only accessible with JWT
	auth := router.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	// Example protected route
	auth.GET("/profile", controller.GetUserProfile)

	// Add more protected routes here
}
