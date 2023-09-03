package entity

import (
	"fmt"
	"hbcase/utils"
)

type Order struct {
	Id          int64  `json:"id"`
	ProductCode string `json:"product_code"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
}

func NewOrder(productCode string, price, quantity int64) *Order {
	return &Order{
		ProductCode: productCode,
		Quantity:    quantity,
		Price:       price,
	}
}

func (order *Order) ValidateOrder() utils.Errors {
	errs := utils.Errors{}

	if order.ProductCode == "" {
		errs.Add("product_code", "order's product code cannot be empty")
	}

	if order.Quantity == 0 {
		errs.Add("quantity", "order's quantity cannot be zero")
	}

	if order.Price <= 0 {
		errs.Add("price", "order's price cannot be smaller than 0")
	}

	return errs
}

func (order *Order) String() string {
	return fmt.Sprintf("Order created; product=%s, quantity=%d", order.ProductCode, order.Quantity)
}

func CalculateOrdersData(orders []Order, target int64) map[string]interface{} {
	data := make(map[string]interface{})
	totalSales := int64(0)

	totalPrice := int64(0)
	for _, order := range orders {
		totalPrice += order.Price
		totalSales += order.Quantity
	}

	data["totalSales"] = len(orders)
	data["turnover"] = target - totalSales

	if len(orders) > 0 {
		data["averageItemPrice"] = totalPrice / int64(len(orders))
	}

	if _, valid := data["averageItemPrice"]; !valid {
		data["averageItemPrice"] = 0
	}
	return data
}
