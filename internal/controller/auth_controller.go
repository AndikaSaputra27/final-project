package controller

import (
	"net/http"

	"github.com/AndikaSaputra27/booking-system/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
}

// NewAuthController creates a new AuthController
func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// Login handles user login requests
func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := c.AuthService.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}
