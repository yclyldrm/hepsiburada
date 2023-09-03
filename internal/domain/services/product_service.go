package services

import (
	"hbcase/internal/domain/entity"
	"hbcase/internal/domain/repository"
)

type ProductService interface {
	SaveProduct(*entity.Product) error
	GetProductInfo(code string) (*entity.Product, error)
}

type productService struct {
	pr repository.ProductRepository
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productService{
		pr: pr,
	}
}

func (s *productService) GetProductInfo(code string) (*entity.Product, error) {
	return s.pr.GetProductInfo(code)
}

func (s *productService) SaveProduct(product *entity.Product) error {
	return s.pr.SaveProduct(product)
}
