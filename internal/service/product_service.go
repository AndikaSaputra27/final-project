package service

import (
	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"github.com/AndikaSaputra27/booking-system/internal/repository"
)

type ProductService interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (p *productService) CreateProduct(product *entity.Product) (*entity.Product, error) {
	return p.productRepo.Create(product)
}

func (p *productService) GetAllProducts() ([]entity.Product, error) {
	return p.productRepo.GetAll()
}
