package controller

import (
	"net/http"

	"github.com/AndikaSaputra27/booking-system/entity"
	"github.com/AndikaSaputra27/booking-system/service"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	ServiceService service.ServiceService
}

func NewServiceController(ss service.ServiceService) *ServiceController {
	return &ServiceController{ServiceService: ss}
}

func (sc *ServiceController) CreateService(c *gin.Context) {
	var svc entity.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := sc.ServiceService.CreateService(&svc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat service"})
		return
	}
	c.JSON(http.StatusCreated, svc)
}
