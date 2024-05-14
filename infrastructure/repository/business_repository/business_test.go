package business_repository

import (
	"payment_gateway_mercadopago/storage"
	"testing"
)

func TestGetAllBusiness(t *testing.T) {
	storage.InitializeTestDB()

	repository := NewRepository(storage.DB)
	err, businesses := repository.GetAllBusiness()
	if err != nil {
		t.Error(err)
	}

	if len(*businesses) != 3 {
		t.Error("expected 3 businesses")
	}
}

func TestGetBusinessByID(t *testing.T) {
	storage.InitializeTestDB()
	repository := NewRepository(storage.DB)
	err, business := repository.GetBusinessByID(1)
	if err != nil {
		t.Error(err)
	}

	if business.ID != 1 {
		t.Error("expected business with id 1")
	}
	if business.Name != "Nexus Store" {
		t.Error("expected business with name Nexus Store")
	}
}
