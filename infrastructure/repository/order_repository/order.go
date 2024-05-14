package order_repository

import (
	"errors"
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/order_model"
	"strconv"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (orderRepository Repository) CreateOrder(order *order_model.Order, details *[]order_model.OrderDetail) error {
	tx := orderRepository.db.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, detail := range *details {
		detail.OrderID = order.ID
		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("failed to commit order transaction")
	}
	return nil
}

func (orderRepository Repository) GetOrderByID(id uint) (error, *order_model.Order) {
	var order *order_model.Order
	if err := orderRepository.db.Preload("Business").Preload("User").Find(&order, id).Error; err != nil {
		return err, nil
	}
	if order.ID == 0 {
		return errors.New("order with ID: " + strconv.Itoa(int(id)) + " not found"), nil
	}
	return nil, order
}

func (orderRepository Repository) SaveMPPreference(order *order_model.Order) error {
	if err := orderRepository.db.Save(&order).Error; err != nil {
		return err
	}
	return nil
}

func (orderRepository Repository) GetOrdersByUser(userID uint) (error, *[]order_model.Order) {
	var orders *[]order_model.Order
	if err := orderRepository.db.Where("user_id = ? ", userID).Find(&orders).Error; err != nil {
		return err, nil
	}
	return nil, orders
}

func (orderRepository Repository) GetOrderDetail(id uint) (error, *[]order_model.OrderDetail) {
	var orderDetails *[]order_model.OrderDetail
	if err := orderRepository.db.Where("order_id = ?", id).Preload("Product").Find(&orderDetails).Error; err != nil {
		return err, nil
	}
	return nil, orderDetails
}
