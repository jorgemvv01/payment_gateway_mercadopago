package order_service

import (
	"payment_gateway_mercadopago/domain/interfaces/order_interface"
	"payment_gateway_mercadopago/domain/models/order_model"
)

type Service struct {
	iOrder order_interface.IOrder
}

func NewService(iOrder order_interface.IOrder) *Service {
	return &Service{
		iOrder: iOrder,
	}
}

func (orderService *Service) CreateOrder(order *order_model.Order, details *[]order_model.OrderDetail) error {
	return orderService.iOrder.CreateOrder(order, details)
}

func (orderService *Service) GetOrderByID(id uint) (error, *order_model.Order) {
	return orderService.iOrder.GetOrderByID(id)
}

func (orderService *Service) SaveMPPreference(order *order_model.Order) error {
	return orderService.iOrder.SaveMPPreference(order)
}

func (orderService *Service) GetOrdersByUser(userID uint) (error, *[]order_model.Order) {
	return orderService.iOrder.GetOrdersByUser(userID)
}

func (orderService *Service) GetOrderDetail(id uint) (error, *[]order_model.OrderDetail) {
	return orderService.iOrder.GetOrderDetail(id)
}
