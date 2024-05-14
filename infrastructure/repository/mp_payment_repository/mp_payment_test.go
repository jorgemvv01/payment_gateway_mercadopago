package mp_payment_repository

import (
	"payment_gateway_mercadopago/domain/models/mp_payment_model"
	"payment_gateway_mercadopago/domain/models/order_model"
	"payment_gateway_mercadopago/infrastructure/repository/order_repository"
	"payment_gateway_mercadopago/storage"
	"testing"
	"time"
)

func TestCreateAndGetMPPaymentByOrder(t *testing.T) {
	storage.InitializeTestDB()

	orderRepository := order_repository.NewRepository(storage.DB)
	order := order_model.Order{
		BusinessID:            1,
		UserID:                1,
		MpClientID:            "",
		MpPreferenceID:        "",
		MpInitPoint:           "",
		MpPreferenceCreatedAt: time.Time{},
		Discount:              10.0,
		Tax:                   19.0,
		Subtotal:              120000.0,
		Total:                 100000.0,
	}
	details := []order_model.OrderDetail{
		{
			ProductID:         1,
			Quantity:          1,
			Price:             1200,
			Discount:          12,
			UnitDiscountValue: 1200,
			Tax:               19,
			UnitTaxValue:      2100,
		},
	}
	err := orderRepository.CreateOrder(&order, &details)

	repository := NewRepository(storage.DB)
	mpPayment := mp_payment_model.MPPayment{
		OrderID:           1,
		PaymentID:         10220102,
		Status:            "pending",
		StatusDetail:      "pending",
		Currency:          "COP",
		TransactionAmount: 100000,
		TaxesAmount:       0,
		Payer:             "{}",
		PaymentMethod:     "{}",
		IPAddress:         "127.0.0.1",
		DateApproved:      time.Time{},
		DateCreated:       time.Time{},
		DateLastUpdated:   time.Time{},
	}
	err = repository.CreateMPPayment(&mpPayment)
	if err != nil {
		t.Error(err)
	}

	if err = storage.DB.Last(&mpPayment).Error; err != nil {
		t.Error(err)
	}
	if mpPayment.TransactionAmount != 100000 {
		t.Error("transaction amount should be 100000")
	}

	err, mpPaymentData := repository.GetMPPaymentByOrder(mpPayment.OrderID)
	if err != nil {
		t.Error(err)
	}

	if mpPaymentData.PaymentID != 10220102 {
		t.Error("payment_id should be 10220102")
	}
}
