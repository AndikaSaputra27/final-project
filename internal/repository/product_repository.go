package repository

import (
	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *entity.Product) (*entity.Product, error)
	GetAll() ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *entity.Product) (*entity.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
