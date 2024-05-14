package business_service

import (
	"payment_gateway_mercadopago/domain/interfaces/business_interface"
	"payment_gateway_mercadopago/domain/models/business_model"
)

type Service struct {
	iBusiness business_interface.IBusiness
}

func NewService(iBusiness business_interface.IBusiness) *Service {
	return &Service{
		iBusiness: iBusiness,
	}
}

func (businessService *Service) GetBusinessByID(id uint) (error, *business_model.Business) {
	return businessService.iBusiness.GetBusinessByID(id)
}

func (businessService *Service) GetAllBusiness() (error, *[]business_model.Business) {
	return businessService.iBusiness.GetAllBusiness()
}
