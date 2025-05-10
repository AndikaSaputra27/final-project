package service

import (
	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"github.com/AndikaSaputra27/booking-system/internal/repository"
)

type ServiceService interface {
	CreateService(service *entity.Service) (*entity.Service, error)
	GetAllServices() ([]entity.Service, error)
}

type serviceService struct {
	serviceRepo repository.ServiceRepository
}

func NewServiceService(serviceRepo repository.ServiceRepository) ServiceService {
	return &serviceService{serviceRepo: serviceRepo}
}

func (s *serviceService) CreateService(service *entity.Service) (*entity.Service, error) {
	return s.serviceRepo.Create(service)
}

func (s *serviceService) GetAllServices() ([]entity.Service, error) {
	return s.serviceRepo.GetAll()
}
