package log_interface

import (
	"payment_gateway_mercadopago/domain/models/log_model"
)

type ILog interface {
	CreateLog(log *log_model.Log) error
}
