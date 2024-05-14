package product_interface

import (
	"payment_gateway_mercadopago/domain/models/product_model"
)

type IProduct interface {
	GetProductsByIDs(ids []uint) (error, *[]product_model.Product)
	GetAllProductsByBusiness(businessID uint) (error, *[]product_model.Product)
	GetPromotionalProductsByBusiness(businessID uint) (error, *[]product_model.Product)
}
