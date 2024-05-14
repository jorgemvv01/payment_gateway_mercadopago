package business_interface

import (
	"payment_gateway_mercadopago/domain/models/business_model"
)

type IBusiness interface {
	GetBusinessByID(id uint) (error, *business_model.Business)
	GetAllBusiness() (error, *[]business_model.Business)
}
