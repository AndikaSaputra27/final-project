package controller

import (
	"net/http"

	"github.com/AndikaSaputra27/booking-system/internal/config"
	"github.com/AndikaSaputra27/booking-system/internal/dto"
	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := entity.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func GetUserProfile(c *gin.Context) {
	username := c.GetString("username")
	role := c.GetString("role")

	c.JSON(http.StatusOK, gin.H{
		"message":  "Profile fetched successfully",
		"username": username,
		"role":     role,
	})
}
