package product_model

import (
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/business_model"
)

type Product struct {
	gorm.Model
	BusinessID  uint `gorm:"not null"`
	Business    business_model.Business
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Code        string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Discount    float32 `gorm:"not null"`
	Tax         float32 `gorm:"not null"`
	Image       string
}

type ProductResponse struct {
	ID          uint
	BusinessID  uint    `gorm:"not null"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Code        string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Discount    float32 `gorm:"not null"`
	Tax         float32 `gorm:"not null"`
	Image       string
}

func NewProductResponse(product Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		BusinessID:  product.BusinessID,
		Name:        product.Name,
		Description: product.Description,
		Code:        product.Code,
		Price:       product.Price,
		Discount:    product.Discount,
		Tax:         product.Tax,
		Image:       product.Image,
	}
}
