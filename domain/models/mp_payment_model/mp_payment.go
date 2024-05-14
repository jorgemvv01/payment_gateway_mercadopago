package mp_payment_model

import (
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/order_model"
	"time"
)

type MPPayment struct {
	gorm.Model
	OrderID           uint `gorm:"not null"`
	Order             order_model.Order
	PaymentID         uint `gorm:"not null"`
	Status            string
	StatusDetail      string
	Currency          string
	TransactionAmount float64
	TaxesAmount       float64
	Payer             string
	PaymentMethod     string
	IPAddress         string
	DateApproved      time.Time
	DateCreated       time.Time
	DateLastUpdated   time.Time
}

type MPPaymentResponse struct {
	ID                uint
	OrderID           uint
	PaymentID         uint
	Status            string
	StatusDetail      string
	Currency          string
	TransactionAmount float64
	TaxesAmount       float64
	Payer             string
	PaymentMethod     string
	IPAddress         string
	DateApproved      time.Time
	DateCreated       time.Time
	DateLastUpdated   time.Time
}

func NewMPPaymentResponse(mpPayment MPPayment) MPPaymentResponse {
	return MPPaymentResponse{
		ID:                mpPayment.ID,
		OrderID:           mpPayment.OrderID,
		PaymentID:         mpPayment.PaymentID,
		Status:            mpPayment.Status,
		StatusDetail:      mpPayment.StatusDetail,
		Currency:          mpPayment.Currency,
		TransactionAmount: mpPayment.TransactionAmount,
		TaxesAmount:       mpPayment.TaxesAmount,
		Payer:             mpPayment.Payer,
		PaymentMethod:     mpPayment.PaymentMethod,
		IPAddress:         mpPayment.IPAddress,
		DateApproved:      mpPayment.DateApproved,
		DateCreated:       mpPayment.DateCreated,
		DateLastUpdated:   mpPayment.DateLastUpdated,
	}
}
