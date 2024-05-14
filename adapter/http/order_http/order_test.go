package order_http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"payment_gateway_mercadopago/domain/models/response_model"
	"payment_gateway_mercadopago/storage"
	"strings"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	storage.InitializeTestDB()
	orderRequestJSON, err := os.Open("../../../mocks/order_mock/create_order_request.json")
	if err != nil {
		t.Fatal(err)
	}
	orderRequest, err := io.ReadAll(orderRequestJSON)
	if err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	controller := NewOrderHTTP()
	router.POST("/api/order", controller.CreateOrder)

	//Ok test
	responseRecorder := httptest.NewRecorder()
	requestBody := orderRequest
	request := httptest.NewRequest("POST", "/api/order", strings.NewReader(string(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	//Failed test
	failedOrderRequestJSON, err := os.Open("../../../mocks/order_mock/failed_create_order_request.json")
	if err != nil {
		t.Fatal(err)
	}
	failedOrderRequest, err := io.ReadAll(failedOrderRequestJSON)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder = httptest.NewRecorder()
	requestBody = failedOrderRequest
	request = httptest.NewRequest("POST", "/api/order", strings.NewReader(string(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	var responseBody response_model.Response
	if err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	if responseBody.Message != "business with ID: 1009 not found" {
		t.Error("no business must exist with ID: 1009")
	}
}

func TestMPFeedbackOrder(t *testing.T) {
	storage.InitializeTestDB()
	feedbackRequestJSON, err := os.Open("../../../mocks/order_mock/feedback_request.json")
	if err != nil {
		t.Fatal(err)
	}
	feedbackRequest, err := io.ReadAll(feedbackRequestJSON)
	if err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	controller := NewOrderHTTP()
	responseRecorder := httptest.NewRecorder()
	router.POST("/api/order/feedback", controller.MPFeedback)

	requestBody := feedbackRequest
	request := httptest.NewRequest("POST", "/api/order/feedback?business_id=1&type=payment&data.id=123456789", strings.NewReader(string(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
