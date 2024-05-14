package mp_payment_repository

import (
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/mp_payment_model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (mpPaymentRepository Repository) CreateMPPayment(payment *mp_payment_model.MPPayment) error {
	return mpPaymentRepository.db.Create(&payment).Error
}

func (mpPaymentRepository Repository) GetMPPaymentByOrder(orderID uint) (error, *mp_payment_model.MPPayment) {
	var mpPayment *mp_payment_model.MPPayment
	if err := mpPaymentRepository.db.Where("order_id = ?", orderID).Order("created_at desc").Find(&mpPayment).Error; err != nil {
		return err, nil
	}
	return nil, mpPayment
}
