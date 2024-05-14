package mp_payment_http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"payment_gateway_mercadopago/domain/models/response_model"
	"payment_gateway_mercadopago/storage"
	"testing"
)

func TestGetMPPaymentByOrder(t *testing.T) {
	storage.InitializeTestDB()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	controller := NewMPPaymentHTTP()
	responseRecorder := httptest.NewRecorder()
	router.GET("/api/mp-payment/:id", controller.GetMPPaymentByOrder)

	request := httptest.NewRequest("GET", "/api/mp-payment/1000", nil)
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	var responseBody response_model.Response
	if err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}

	if responseBody.Message != "no mp payment found" {
		t.Error("mp payment should not have been found")
	}
}
