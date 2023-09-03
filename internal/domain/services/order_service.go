package services

import (
	"hbcase/internal/domain/entity"
	"hbcase/internal/domain/repository"
)

type OrderService interface {
	Create(*entity.Order) error
	GetOrders(productCode string) ([]entity.Order, error)
}

type orderService struct {
	or repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		or: orderRepo,
	}
}

func (os *orderService) Create(orderEntity *entity.Order) error {
	return os.or.Create(orderEntity)
}

func (os *orderService) GetOrders(productCode string) ([]entity.Order, error) {
	return os.or.GetOrders(productCode)
}
