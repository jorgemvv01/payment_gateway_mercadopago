package product_http

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

func TestGetAllProductsByBusiness(t *testing.T) {
	storage.InitializeTestDB()
	productsResponseJSON, err := os.Open("../../../mocks/product_mock/products_by_business_response.json")
	if err != nil {
		t.Fatal(err)
	}
	productsResponse, err := io.ReadAll(productsResponseJSON)
	if err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	controller := NewProductHTTP()
	responseRecorder := httptest.NewRecorder()
	router.GET("/api/product/by-business/:id", controller.GetAllProductsByBusiness)

	request := httptest.NewRequest("GET", "/api/product/by-business/1", nil)
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody response_model.Response
	if err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	if responseBody.Message != "products found" {
		t.Error("products should have been found")
	}

	data, err := json.Marshal(responseBody.Data)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, productsResponse) {
		t.Error("no match of the products obtained with the mock")
	}
}

func TestGetPromotionalProductsByBusiness(t *testing.T) {
	storage.InitializeTestDB()
	productsResponseJSON, err := os.Open("../../../mocks/product_mock/promotional_products_response.json")
	if err != nil {
		t.Fatal(err)
	}
	productsResponse, err := io.ReadAll(productsResponseJSON)
	if err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	controller := NewProductHTTP()
	responseRecorder := httptest.NewRecorder()
	router.GET("/api/product/promotional/:id", controller.GetPromotionalProductsByBusiness)

	request := httptest.NewRequest("GET", "/api/product/promotional/1", nil)
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody response_model.Response
	if err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	if responseBody.Message != "products found" {
		t.Error("products should have been found")
	}

	data, err := json.Marshal(responseBody.Data)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, productsResponse) {
		t.Error("no match of the products obtained with the mock")
	}
}
