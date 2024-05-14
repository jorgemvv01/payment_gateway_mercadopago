package mp_payment_interface

import (
	"payment_gateway_mercadopago/domain/models/mp_payment_model"
)

type IMPPayment interface {
	CreateMPPayment(payment *mp_payment_model.MPPayment) error
	GetMPPaymentByOrder(orderID uint) (error, *mp_payment_model.MPPayment)
}
