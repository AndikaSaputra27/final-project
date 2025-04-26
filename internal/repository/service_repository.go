package repository

import (
	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(service *entity.Service) error
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db}
}

func (r *serviceRepository) Create(service *entity.Service) error {
	return r.db.Create(service).Error
}
