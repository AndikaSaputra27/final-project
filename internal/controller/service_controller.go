package controller

import (
	"net/http"

	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"github.com/AndikaSaputra27/booking-system/internal/service"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	serviceService service.ServiceService
}

func NewServiceController(serviceService service.ServiceService) *ServiceController {
	return &ServiceController{serviceService: serviceService}
}

func (c *ServiceController) CreateService(ctx *gin.Context) {
	var service entity.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newService, err := c.serviceService.CreateService(&service)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newService)
}

func (c *ServiceController) GetAllServices(ctx *gin.Context) {
	services, err := c.serviceService.GetAllServices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, services)
}
