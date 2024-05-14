package user_repository

import (
	"payment_gateway_mercadopago/storage"
	"testing"
)

func TestGetUserByID(t *testing.T) {
	storage.InitializeTestDB()

	repository := NewRepository(storage.DB)
	err, user := repository.GetUserByID(1)
	if err != nil {
		t.Error(err)
	}

	if user.ID != 1 {
		t.Error("user id should be 1")
	}

	if user.Name != "John" {
		t.Error("user name should be 'John'")
	}

	if user.LastName != "Doe" {
		t.Error("user last name should be 'Doe'")
	}
}
