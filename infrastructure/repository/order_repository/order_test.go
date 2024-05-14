package order_repository

import (
	"payment_gateway_mercadopago/domain/models/order_model"
	"payment_gateway_mercadopago/storage"
	"testing"
	"time"
)

func TestCreateAndGetOrders(t *testing.T) {
	storage.InitializeTestDB()

	repository := NewRepository(storage.DB)
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
	err := repository.CreateOrder(&order, &details)
	if err != nil {
		t.Error(err)
	}
	if err = storage.DB.Last(&order).Error; err != nil {
		t.Error(err)
	}

	if order.Subtotal != 120000 {
		t.Error("Subtotal should be 120000")
	}

	order.MpInitPoint = "https://mercadopago.com"
	order.MpClientID = "102010202"
	order.MpPreferenceID = "02sjd-12mord-1mrt-59mdma"
	order.MpPreferenceCreatedAt = time.Time{}
	err = repository.SaveMPPreference(&order)
	if err != nil {
		t.Error(err)
	}

	err, orderByID := repository.GetOrderByID(order.ID)
	if err != nil {
		t.Error(err)
	}

	if orderByID.Tax != 19.0 {
		t.Error("order tax should be 19.0")
	}

	err, ordersByUser := repository.GetOrdersByUser(1)
	if err != nil {
		t.Error(err)
	}

	if len(*ordersByUser) == 0 {
		t.Error("orders by user should not be empty")
	}

	lastPosition := len(*ordersByUser) - 1
	if (*ordersByUser)[lastPosition].UserID != 1 {
		t.Error("order user id should be 1")
	}

	if (*ordersByUser)[lastPosition].Subtotal != 120000.0 {
		t.Error("order subtotal should be 120000.0")
	}

	if (*ordersByUser)[lastPosition].MpInitPoint != "https://mercadopago.com" {
		t.Error("order mp_init_point should be https://mercadopago.com")
	}

	err, orderDetails := repository.GetOrderDetail(order.ID)
	if err != nil {
		t.Error(err)
	}

	if len(*orderDetails) != 1 {
		t.Error("orderDetails should be 1")
	}

	if (*orderDetails)[0].ProductID != 1 {
		t.Error("order product_id should be 1")
	}
}
