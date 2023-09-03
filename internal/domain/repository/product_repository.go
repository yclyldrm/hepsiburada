package repository

import (
	"hbcase/internal/domain/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	SaveProduct(*entity.Product) error
	GetProductInfo(code string) (*entity.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}

func (r *productRepository) SaveProduct(product *entity.Product) error {
	return r.DB.Save(product).Error
}

func (r *productRepository) GetProductInfo(code string) (*entity.Product, error) {
	var product entity.Product
	err := r.DB.Where("product_code = ?", code).Order("-id").First(&product).Error
	return &product, err
}
