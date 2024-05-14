package product_service

import (
	"payment_gateway_mercadopago/domain/interfaces/product_interface"
	"payment_gateway_mercadopago/domain/models/product_model"
)

type Service struct {
	iProduct product_interface.IProduct
}

func NewService(iProduct product_interface.IProduct) *Service {
	return &Service{
		iProduct: iProduct,
	}
}

func (productService *Service) GetProductsByIDs(ids []uint) (error, *[]product_model.Product) {
	return productService.iProduct.GetProductsByIDs(ids)
}

func (productService *Service) GetAllProductsByBusiness(businessID uint) (error, *[]product_model.Product) {
	return productService.iProduct.GetAllProductsByBusiness(businessID)
}

func (productService *Service) GetPromotionalProductsByBusiness(businessID uint) (error, *[]product_model.Product) {
	return productService.iProduct.GetPromotionalProductsByBusiness(businessID)
}
