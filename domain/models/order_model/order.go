package order_model

import (
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/business_model"
	"payment_gateway_mercadopago/domain/models/user_model"
	"time"
)

type Order struct {
	gorm.Model
	BusinessID            uint `gorm:"not null"`
	Business              business_model.Business
	UserID                uint `gorm:"not null"`
	User                  user_model.User
	MpClientID            string
	MpPreferenceID        string
	MpInitPoint           string
	MpPreferenceCreatedAt time.Time
	Discount              float64 `gorm:"not null"`
	Tax                   float64 `gorm:"not null"`
	Subtotal              float64 `gorm:"not null"`
	Total                 float64 `gorm:"not null"`
}

type OrderRequest struct {
	BusinessID uint             `json:"business_id" binding:"required"`
	UserID     uint             `json:"user_id" binding:"required"`
	Products   []ProductRequest `json:"products" binding:"required"`
}

type ProductRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

type OrderResponse struct {
	ID                    uint
	BusinessID            uint
	UserID                uint
	MpClientID            string
	MpPreferenceID        string
	MpInitPoint           string
	MpPreferenceCreatedAt time.Time
	Discount              float64
	Tax                   float64
	Subtotal              float64
	Total                 float64
}

func NewOrderResponse(order Order) OrderResponse {
	return OrderResponse{
		ID:                    order.ID,
		BusinessID:            order.BusinessID,
		UserID:                order.UserID,
		MpClientID:            order.MpClientID,
		MpPreferenceID:        order.MpPreferenceID,
		MpInitPoint:           order.MpInitPoint,
		MpPreferenceCreatedAt: order.MpPreferenceCreatedAt,
		Discount:              order.Discount,
		Tax:                   order.Tax,
		Subtotal:              order.Subtotal,
		Total:                 order.Total,
	}
}
