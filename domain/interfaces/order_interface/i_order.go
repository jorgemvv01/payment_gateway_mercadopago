package order_interface

import (
	"payment_gateway_mercadopago/domain/models/order_model"
)

type IOrder interface {
	CreateOrder(order *order_model.Order, details *[]order_model.OrderDetail) error
	GetOrderByID(id uint) (error, *order_model.Order)
	GetOrdersByUser(userID uint) (error, *[]order_model.Order)
	SaveMPPreference(order *order_model.Order) error
	GetOrderDetail(id uint) (error, *[]order_model.OrderDetail)
}
