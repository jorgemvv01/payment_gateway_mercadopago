package product_repository

import (
	"payment_gateway_mercadopago/storage"
	"testing"
)

func TestGetProductsByIDs(t *testing.T) {
	storage.InitializeTestDB()

	repository := NewRepository(storage.DB)
	ids := []uint{1, 2}
	err, products := repository.GetProductsByIDs(ids)
	if err != nil {
		t.Error(err)
	}

	if len(*products) != 2 {
		t.Error("expected 2 products")
	}
}

func TestGetAllProductsByBusiness(t *testing.T) {
	storage.InitializeTestDB()

	repository := NewRepository(storage.DB)
	err, products := repository.GetAllProductsByBusiness(1)
	if err != nil {
		t.Error(err)
	}

	if len(*products) != 5 {
		t.Error("expected 5 products")
	}
}

func TestGetPromotionalProductsByBusiness(t *testing.T) {
	storage.InitializeTestDB()

	repository := NewRepository(storage.DB)
	err, products := repository.GetPromotionalProductsByBusiness(1)
	if err != nil {
		t.Error(err)
	}

	if len(*products) != 4 {
		t.Error("expected 4 products")
	}
}
