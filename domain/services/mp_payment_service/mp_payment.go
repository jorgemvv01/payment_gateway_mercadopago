package mp_payment_service

import (
	"payment_gateway_mercadopago/domain/interfaces/mp_payment_interface"
	"payment_gateway_mercadopago/domain/models/mp_payment_model"
)

type Service struct {
	iMPPayment mp_payment_interface.IMPPayment
}

func NewService(iMPPayment mp_payment_interface.IMPPayment) *Service {
	return &Service{
		iMPPayment: iMPPayment,
	}
}

func (mpPaymentService *Service) CreateMPPayment(payment *mp_payment_model.MPPayment) error {
	return mpPaymentService.iMPPayment.CreateMPPayment(payment)
}

func (mpPaymentService *Service) GetMPPaymentByOrder(orderID uint) (error, *mp_payment_model.MPPayment) {
	return mpPaymentService.iMPPayment.GetMPPaymentByOrder(orderID)
}
