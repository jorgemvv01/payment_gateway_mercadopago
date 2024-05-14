package log_repository

import (
	"payment_gateway_mercadopago/domain/models/log_model"
	"payment_gateway_mercadopago/storage"
	"testing"
)

func TestCreateLog(t *testing.T) {
	storage.InitializeTestDB()

	repository := NewRepository(storage.DB)
	log := log_model.Log{
		LogType:     "error",
		Status:      "failed",
		Information: "error log test",
		Details:     "error log test details",
		Module:      "log",
	}
	err := repository.CreateLog(&log)
	if err != nil {
		t.Error(err)
	}
	if err = storage.DB.Last(&log).Error; err != nil {
		t.Error(err)
	}

	if log.Status != "failed" {
		t.Error("status should be failed")
	}
	if log.Information != "error log test" {
		t.Error("information should be error log test")
	}
	if log.Details != "error log test details" {
		t.Error("details should be error log test details")
	}
	if log.Module != "log" {
		t.Error("module should be log")
	}
}
