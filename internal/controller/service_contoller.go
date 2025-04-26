package controller

import (
	"net/http"

	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"github.com/AndikaSaputra27/booking-system/internal/service"
	"github.com/gin-gonic/gin"
)

type BookingServiceController struct {
	BookingService service.BookingService
}

func NewBookingServiceController(bs service.BookingService) *BookingServiceController {
	return &BookingServiceController{BookingService: bs}
}

func (sc *BookingServiceController) CreateService(c *gin.Context) {
	var svc entity.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := sc.BookingService.CreateService(&svc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat service"})
		return
	}
	c.JSON(http.StatusCreated, svc)
}
