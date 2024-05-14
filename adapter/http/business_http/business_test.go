package business_http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"payment_gateway_mercadopago/domain/models/response_model"
	"payment_gateway_mercadopago/storage"
	"testing"
)

func TestGetAllBusiness(t *testing.T) {
	storage.InitializeTestDB()
	businessResponseJSON, err := os.Open("../../../mocks/business_mock/business_response.json")
	if err != nil {
		t.Fatal(err)
	}
	businessResponse, err := io.ReadAll(businessResponseJSON)
	if err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	controller := NewBusinessHTTP()
	responseRecorder := httptest.NewRecorder()
	router.GET("/api/business", controller.GetAllBusiness)

	request := httptest.NewRequest("GET", "/api/business", nil)
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody response_model.Response
	if err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	if responseBody.Message != "business found" {
		t.Error("businesses should have been found")
	}

	data, err := json.Marshal(responseBody.Data)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, businessResponse) {
		t.Error("no match of the business obtained with the mock")
	}
}
