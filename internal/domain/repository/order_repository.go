package repository

import (
	"hbcase/internal/domain/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *entity.Order) error
	GetOrders(productCode string) ([]entity.Order, error)
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (r *orderRepository) Create(order *entity.Order) error {
	return r.DB.Save(order).Error
}

func (r *orderRepository) GetOrders(productCode string) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.DB.Where("product_code = ?", productCode).Find(&orders).Error
	return orders, err
}
