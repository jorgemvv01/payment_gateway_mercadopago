package order_http

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/pkg/preference"
	"net/http"
	"payment_gateway_mercadopago/domain/models/log_model"
	"payment_gateway_mercadopago/domain/models/order_model"
	"payment_gateway_mercadopago/domain/models/response_model"
	"payment_gateway_mercadopago/domain/services/business_service"
	"payment_gateway_mercadopago/domain/services/log_service"
	"payment_gateway_mercadopago/domain/services/mp_payment_service"
	"payment_gateway_mercadopago/domain/services/order_service"
	"payment_gateway_mercadopago/infrastructure/repository/business_repository"
	"payment_gateway_mercadopago/infrastructure/repository/log_repository"
	"payment_gateway_mercadopago/infrastructure/repository/mp_payment_repository"
	"payment_gateway_mercadopago/infrastructure/repository/order_repository"
	"payment_gateway_mercadopago/storage"
	"strconv"
)

type OrderHTTP struct{}

func NewOrderHTTP() *OrderHTTP {
	return &OrderHTTP{}
}

// CreateOrder
// @Summary Create order with Mercado Pago
// @Description Create order with Mercado Pago.
// @Param tags body order_model.OrderRequest true "The following body is required"
// @Produce application/json
// @Tags Order
// @Success 201 {object} response_model.Response{}
// @Failure 400 {object} response_model.Response{}
// @Failure 500 {object} response_model.Response{}
// @Router /order [post]
func (controller *OrderHTTP) CreateOrder(c *gin.Context) {
	var err, orderRequest = controller.validateOrderRequestForm(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	err, products := controller.validateOrderRequestData(orderRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	order, details := controller.calculateOrder(orderRequest, products)
	orderRepository := order_repository.NewRepository(storage.DB)
	orderService := order_service.NewService(orderRepository)

	if err = orderService.CreateOrder(order, details); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response_model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	err, order = orderService.GetOrderByID(order.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusCreated, response_model.Response{
			Status:  "warning",
			Message: `order created without preference ` + err.Error(),
			Data:    order_model.NewOrderResponse(*order),
		})
		return
	}

	err, cfg, request := controller.createPreferenceRequest(*order, *details, *products)
	client := preference.NewClient(cfg)
	resource, err := client.Create(c, *request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusCreated, response_model.Response{
			Status:  "warning",
			Message: `order created without preference ` + err.Error(),
			Data:    order_model.NewOrderResponse(*order),
		})
		return
	}
	controller.updateMPPreferenceOrder(order, *resource)
	if err = orderService.SaveMPPreference(order); err != nil {
		c.AbortWithStatusJSON(http.StatusCreated, response_model.Response{
			Status:  "warning",
			Message: `order created without preference ` + err.Error(),
			Data:    order_model.NewOrderResponse(*order),
		})
		return
	}

	c.JSON(http.StatusCreated, response_model.Response{
		Status:  "success",
		Message: "order created successfully",
		Data:    order_model.NewOrderResponse(*order),
	})

}

func (controller *OrderHTTP) MPFeedback(c *gin.Context) {
	_type, exists := c.GetQuery("type")
	if !exists {
		c.Status(http.StatusNoContent)
		return
	}

	if _type != "payment" {
		c.Status(http.StatusNoContent)
		return
	}

	logRepository := log_repository.NewRepository(storage.DB)
	logService := log_service.NewService(logRepository)
	err, query, body := controller.getMPPaymentConfirmationData(c)
	if err != nil {
		logData := log_model.Log{
			LogType:     "error",
			Status:      "failed",
			Information: "there was an error obtaining the query and body",
			Details:     err.Error(),
			Module:      "mercado pago - payment confirmation",
		}
		_ = logService.CreateLog(&logData)
		c.Status(http.StatusInternalServerError)
		return
	}

	businessID, exists := c.GetQuery("business_id")
	if !exists {
		logData := log_model.Log{
			LogType:     "error",
			Status:      "failed",
			Information: "no business ID found in query",
			Details:     `query received: ` + string(query) + ` body received: ` + string(body),
			Module:      "mercado pago - payment confirmation",
		}
		_ = logService.CreateLog(&logData)
		c.Status(http.StatusInternalServerError)
		return
	}

	businessRepository := business_repository.NewRepository(storage.DB)
	businessService := business_service.NewService(businessRepository)
	businessIDUINT, err := strconv.ParseUint(businessID, 10, 64)
	if err != nil {
		logData := log_model.Log{
			LogType:     "error",
			Status:      "failed",
			Information: "business id could not be converted to integer",
			Details:     `query received: ` + string(query) + ` body received: ` + string(body),
			Module:      "mercado pago - payment confirmation",
		}
		_ = logService.CreateLog(&logData)
		c.Status(http.StatusInternalServerError)
		return
	}
	err, business := businessService.GetBusinessByID(uint(businessIDUINT))
	if err != nil {
		logData := log_model.Log{
			LogType:     "error",
			Status:      "failed",
			Information: "business not found in database",
			Details:     `query received: ` + string(query) + ` body received: ` + string(body),
			Module:      "mercado pago - payment confirmation",
		}
		_ = logService.CreateLog(&logData)
		c.Status(http.StatusInternalServerError)
		return
	}

	cfg, err := config.New(business.MpToken)
	if err != nil {
		logData := log_model.Log{
			LogType:     "error",
			Status:      "failed",
			Information: "business not found in database",
			Details:     `query received: ` + string(query) + ` body received: ` + string(body),
			Module:      "mercado pago - payment confirmation",
		}
		_ = logService.CreateLog(&logData)
		c.Status(http.StatusInternalServerError)
		return
	}

	id := c.Query("data.id")
	paymentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logData := log_model.Log{
			LogType:     "error",
			Status:      "failed",
			Information: "mp order id could not be converted to integer",
			Details:     `query received: ` + string(query) + ` body received: ` + string(body),
			Module:      "mercado pago - payment confirmation",
		}
		_ = logService.CreateLog(&logData)
		c.Status(http.StatusInternalServerError)
		return
	}

	client := payment.NewClient(cfg)
	paymentData, err := client.Get(c, int(paymentID))
	if err != nil {
		logData := log_model.Log{
			LogType:     "error",
			Status:      "failed",
			Information: "mp payment not found",
			Details:     `query received: ` + string(query) + ` body received: ` + string(body),
			Module:      "mercado pago - payment confirmation",
		}
		_ = logService.CreateLog(&logData)
		c.Status(http.StatusInternalServerError)
		return
	}

	err, mpPayment := controller.createMPPayment(*paymentData)
	mpPaymentRepository := mp_payment_repository.NewRepository(storage.DB)
	mpPaymentService := mp_payment_service.NewService(mpPaymentRepository)
	err = mpPaymentService.CreateMPPayment(mpPayment)
	if err != nil {
		logData := log_model.Log{
			LogType:     "error",
			Status:      "failed",
			Information: `mp payment could not be created` + err.Error(),
			Details:     `query received: ` + string(query) + ` body received: ` + string(body),
			Module:      "mercado pago - payment confirmation",
		}
		_ = logService.CreateLog(&logData)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

// GetOrdersByUser
// @Summary Get orders by user
// @Description Get orders by user.
// @Param user_id path string true "The user ID is required in the query"
// @Produce application/json
// @Tags Order
// @Success 200 {object} response_model.Response{}
// @Failure 400 {object} response_model.Response{}
// @Failure 404 {object} response_model.Response{}
// @Failure 500 {object} response_model.Response{}
// @Router /order/user/{user_id} [get]
func (controller *OrderHTTP) GetOrdersByUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "invalid user ID",
		})
		return
	}
	orderRepository := order_repository.NewRepository(storage.DB)
	orderService := order_service.NewService(orderRepository)
	err, orders := orderService.GetOrdersByUser(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response_model.Response{
			Status:  "error",
			Message: `unable to get orders, ` + err.Error(),
		})
		return
	}

	if len(*orders) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, response_model.Response{
			Status:  "error",
			Message: `no orders were found for the user with ID: ` + c.Param("id"),
		})
		return
	}

	var ordersResponse []order_model.OrderResponse
	for _, order := range *orders {
		ordersResponse = append(ordersResponse, order_model.NewOrderResponse(order))
	}

	c.JSON(http.StatusOK, response_model.Response{
		Status:  "success",
		Message: "orders found",
		Data:    ordersResponse,
	})
}

// GetOrderDetail
// @Summary Get order detail
// @Description Get order detail.
// @Param order_id path string true "The order ID is required in the query"
// @Produce application/json
// @Tags Order
// @Success 200 {object} response_model.Response{}
// @Failure 400 {object} response_model.Response{}
// @Failure 404 {object} response_model.Response{}
// @Failure 500 {object} response_model.Response{}
// @Router /order/details/{order_id} [get]
func (controller *OrderHTTP) GetOrderDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "invalid order ID",
		})
		return
	}
	orderRepository := order_repository.NewRepository(storage.DB)
	orderService := order_service.NewService(orderRepository)
	err, orderDetails := orderService.GetOrderDetail(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response_model.Response{
			Status:  "error",
			Message: `unable to get orders, ` + err.Error(),
		})
		return
	}

	if len(*orderDetails) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, response_model.Response{
			Status:  "error",
			Message: `no order details were found for the order with ID: ` + c.Param("id"),
		})
		return
	}

	var orderDetailsResponse []order_model.OrderDetailResponse
	for _, orderDetail := range *orderDetails {
		orderDetailsResponse = append(orderDetailsResponse, order_model.NewOrderDetailResponse(orderDetail))
	}

	c.JSON(http.StatusOK, response_model.Response{
		Status:  "success",
		Message: "order details found",
		Data:    orderDetailsResponse,
	})
}
