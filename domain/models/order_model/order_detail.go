package order_model

import (
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/product_model"
)

type OrderDetail struct {
	gorm.Model
	OrderID           uint `gorm:"not null"`
	Order             Order
	ProductID         uint `gorm:"not null"`
	Product           product_model.Product
	Quantity          uint    `gorm:"not null"`
	Price             float64 `gorm:"not null"`
	Discount          float32 `gorm:"not null"`
	UnitDiscountValue float64 `gorm:"not null"`
	Tax               float32 `gorm:"not null"`
	UnitTaxValue      float64 `gorm:"not null"`
}

type OrderDetailResponse struct {
	ID                uint
	OrderID           uint
	Product           product_model.ProductResponse
	Quantity          uint
	Price             float64
	Discount          float32
	UnitDiscountValue float64
	Tax               float32
	UnitTaxValue      float64
}

func NewOrderDetailResponse(orderDetail OrderDetail) OrderDetailResponse {
	return OrderDetailResponse{
		ID:                orderDetail.ID,
		OrderID:           orderDetail.OrderID,
		Product:           product_model.NewProductResponse(orderDetail.Product),
		Quantity:          orderDetail.Quantity,
		Price:             orderDetail.Price,
		Discount:          orderDetail.Discount,
		UnitDiscountValue: orderDetail.UnitDiscountValue,
		Tax:               orderDetail.Tax,
		UnitTaxValue:      orderDetail.UnitTaxValue,
	}
}
