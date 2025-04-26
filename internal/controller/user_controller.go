package controller

import (
	"net/http"

	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"github.com/AndikaSaputra27/booking-system/internal/service"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	BookingService service.BookingService
}

func NewServiceController(ss service.BookingService) *ServiceController {
	return &ServiceController{BookingService: ss}
}

func (sc *ServiceController) CreateService(c *gin.Context) {
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
