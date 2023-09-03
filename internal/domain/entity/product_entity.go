package entity

import (
	"fmt"
	"hbcase/utils"
)

type Product struct {
	ID          int64  `json:"id"`
	ProductCode string `json:"product_code"`
	Price       int64  `json:"price"`
	FirstPrice  int64  `json:"first_price"`
	Stock       int64  `json:"stock"`
}

func NewProduct(code string, price, stock int64) *Product {
	return &Product{
		ProductCode: code,
		Price:       price,
		FirstPrice:  price,
		Stock:       stock,
	}
}

func (p *Product) String(funcName string) string {
	var str string

	switch funcName {
	case "create_product":
		str = fmt.Sprintf("Product created; code=%s; price=%d; stock=%d", p.ProductCode, p.Price, p.Stock)
	case "get_product_info":
		str = fmt.Sprintf("Product %s info; price=%d, stock=%d", p.ProductCode, p.Price, p.Stock)
	default:
		str = "Unknown function name"
	}

	return str
}

func (p *Product) ChangeProductPrice(duration, limit int64) {
	change := limit / (duration - 1)

	lastPrice := p.Price - change

	if lastPrice > 0 && p.FirstPrice-limit <= lastPrice {
		p.Price = lastPrice
	}
}

func (p *Product) SetPriceatFirst() {
	p.Price = p.FirstPrice
}

func (p *Product) ValidateProduct() *utils.Errors {
	valErrors := utils.NewErrors()
	if p.ProductCode == "" {
		valErrors.Add("product_code", "Product code cannot be empty")
	}
	if p.Price < 0 {
		valErrors.Add("price", "Price cannot be negative")
	}
	if p.Stock <= 0 {
		valErrors.Add("stock", "Stock cannot be negative")
	}

	return valErrors
}

func (p *Product) DecreaseStock(quantity int64) {
	p.Stock = p.Stock - quantity
}
