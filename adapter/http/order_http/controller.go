package order_http

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/pkg/preference"
	"math"
	"os"
	"payment_gateway_mercadopago/domain/models/mp_payment_model"
	"payment_gateway_mercadopago/domain/models/order_model"
	"payment_gateway_mercadopago/domain/models/product_model"
	"payment_gateway_mercadopago/domain/services/business_service"
	"payment_gateway_mercadopago/domain/services/product_service"
	"payment_gateway_mercadopago/domain/services/user_service"
	"payment_gateway_mercadopago/infrastructure/repository/business_repository"
	"payment_gateway_mercadopago/infrastructure/repository/product_repository"
	"payment_gateway_mercadopago/infrastructure/repository/user_repository"
	"payment_gateway_mercadopago/storage"
	"strconv"
)

func (controller *OrderHTTP) validateOrderRequestForm(c *gin.Context) (error, *order_model.OrderRequest) {
	var orderRequest *order_model.OrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		return errors.New(`invalid request body... ` + err.Error()), nil
	}

	if len(orderRequest.Products) == 0 {
		return errors.New("products is empty"), nil
	}

	for _, p := range orderRequest.Products {
		if p.ProductID == 0 || p.Quantity == 0 {
			return errors.New("all products must have an product_id and a quantity valid"), nil
		}
	}
	return nil, orderRequest
}

func (controller *OrderHTTP) validateOrderRequestData(orderRequest *order_model.OrderRequest) (error, *[]product_model.Product) {
	userRepository := user_repository.NewRepository(storage.DB)
	userService := user_service.NewService(userRepository)
	err, _ := userService.GetUserByID(orderRequest.UserID)
	if err != nil {
		return err, nil
	}
	businessRepository := business_repository.NewRepository(storage.DB)
	businessService := business_service.NewService(businessRepository)
	err, _ = businessService.GetBusinessByID(orderRequest.BusinessID)
	if err != nil {
		return err, nil
	}

	productRepository := product_repository.NewRepository(storage.DB)
	productService := product_service.NewService(productRepository)
	var ids []uint

	for _, product := range orderRequest.Products {
		ids = append(ids, product.ProductID)
	}

	err, products := productService.GetProductsByIDs(ids)
	if err != nil {
		return err, nil
	}
	if len(*products) == 0 {
		return errors.New("no products found"), nil
	}

	if len(*products) != len(orderRequest.Products) {
		return errors.New("not all products were found in the database"), nil
	}

	return nil, products
}

func (controller *OrderHTTP) calculateOrder(orderRequest *order_model.OrderRequest, products *[]product_model.Product) (*order_model.Order, *[]order_model.OrderDetail) {
	var subtotal float64
	var tax float64
	var discount float64
	var total float64
	var orderDetails []order_model.OrderDetail
	for index, product := range *products {
		_price := product.Price
		_quantity := orderRequest.Products[index].Quantity
		_tax := product.Tax
		_discount := product.Discount
		var _productDiscount float64 = 0
		subtotal += _price * float64(_quantity)
		if _discount > 0 {
			_productDiscount = (float64(_discount) / 100) * (_price)
			discount += _productDiscount * float64(_quantity)
		}

		var _productTax float64 = 0
		if _tax > 0 {
			_productTax = (float64(_tax) / 100) * (_price - _productDiscount)
			tax += _productTax * float64(_quantity)
			subtotal += _productTax * float64(_quantity)
		}

		orderDetails = append(orderDetails, order_model.OrderDetail{
			ProductID:         product.ID,
			Discount:          _discount,
			UnitDiscountValue: _productDiscount,
			Quantity:          _quantity,
			Price:             _price,
			Tax:               _tax,
			UnitTaxValue:      _productTax,
		})
	}
	total = subtotal - discount
	order := order_model.Order{
		BusinessID: orderRequest.BusinessID,
		UserID:     orderRequest.UserID,
		Subtotal:   subtotal,
		Discount:   discount,
		Tax:        tax,
		Total:      total,
	}
	return &order, &orderDetails
}

func (controller *OrderHTTP) createPreferenceRequest(order order_model.Order, details []order_model.OrderDetail, products []product_model.Product) (error, *config.Config, *preference.Request) {
	cfg, err := config.New(order.Business.MpToken)
	if err != nil {
		return err, nil, nil
	}
	var items []preference.ItemRequest
	for index, detail := range details {
		productList := products
		items = append(items, preference.ItemRequest{
			Title:      productList[index].Name,
			Quantity:   int(detail.Quantity),
			UnitPrice:  math.Round((detail.Price + detail.UnitTaxValue) - detail.UnitDiscountValue),
			CurrencyID: "COP",
		})
	}
	metadata := map[string]any{
		"order_id": order.ID,
	}
	request := preference.Request{
		Items:           items,
		Metadata:        metadata,
		NotificationURL: os.Getenv("MP_FEEDBACK") + `?business_id=` + strconv.Itoa(int(order.BusinessID)) + `&`,
	}
	return nil, cfg, &request
}

func (controller *OrderHTTP) updateMPPreferenceOrder(order *order_model.Order, preference preference.Response) {
	order.MpClientID = preference.ClientID
	order.MpInitPoint = preference.InitPoint
	order.MpPreferenceID = preference.ID
	order.MpPreferenceCreatedAt = preference.DateCreated
}

func (controller *OrderHTTP) getDataMPPaymentConfirmation(c *gin.Context) (error, []byte, []byte) {
	query, err := json.Marshal(c.Request.URL.Query())
	if err != nil {
		return err, nil, nil
	}
	bodyMap := make(map[string]interface{})
	if err = c.Bind(&bodyMap); err != nil {
		return err, nil, nil
	}
	body, err := json.Marshal(bodyMap)
	if err != nil {
		return err, nil, nil
	}
	return nil, query, body
}

func (controller *OrderHTTP) createMPPayment(paymentData payment.Response) (error, *mp_payment_model.MPPayment) {
	orderID := paymentData.Metadata["order_id"].(float64)
	payer, err := json.Marshal(paymentData.Payer)
	if err != nil {
		return err, nil
	}
	paymentMethod, err := json.Marshal(paymentData.PaymentMethod)
	if err != nil {
		return err, nil
	}
	mpPayment := &mp_payment_model.MPPayment{
		OrderID:           uint(orderID),
		PaymentID:         uint(paymentData.ID),
		Status:            paymentData.Status,
		StatusDetail:      paymentData.StatusDetail,
		Currency:          paymentData.CurrencyID,
		TransactionAmount: paymentData.TransactionAmount,
		TaxesAmount:       paymentData.TaxesAmount,
		Payer:             string(payer),
		PaymentMethod:     string(paymentMethod),
		IPAddress:         paymentData.AdditionalInfo.IPAddress,
		DateApproved:      paymentData.DateApproved,
		DateCreated:       paymentData.DateCreated,
		DateLastUpdated:   paymentData.DateLastUpdated,
	}
	return nil, mpPayment
}
