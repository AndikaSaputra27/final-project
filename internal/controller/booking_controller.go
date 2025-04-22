package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	// tambahkan dependency kalau ada, misalnya service
}

func NewBookingController() *BookingController {
	return &BookingController{}
}

// Tambahkan ini:
func (bc *BookingController) GetBookingHandler(c *gin.Context) {
	// sementara untuk testing
	c.JSON(http.StatusOK, gin.H{
		"message": "Get booking success!",
	})
}
