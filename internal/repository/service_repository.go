package repository

import (
	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(service *entity.Service) (*entity.Service, error)
	GetAll() ([]entity.Service, error)
}

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (r *serviceRepository) Create(service *entity.Service) (*entity.Service, error) {
	if err := r.db.Create(service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

func (r *serviceRepository) GetAll() ([]entity.Service, error) {
	var services []entity.Service
	if err := r.db.Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}
