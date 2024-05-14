package log_service

import (
	"payment_gateway_mercadopago/domain/interfaces/log_interface"
	"payment_gateway_mercadopago/domain/models/log_model"
)

type Service struct {
	iLog log_interface.ILog
}

func NewService(iLog log_interface.ILog) *Service {
	return &Service{
		iLog: iLog,
	}
}

func (logService *Service) CreateLog(log *log_model.Log) error {
	return logService.iLog.CreateLog(log)
}
