package user_service

import (
	"payment_gateway_mercadopago/domain/interfaces/user_interface"
	"payment_gateway_mercadopago/domain/models/user_model"
)

type Service struct {
	iUser user_interface.IUser
}

func NewService(iUser user_interface.IUser) *Service {
	return &Service{
		iUser: iUser,
	}
}

func (userService *Service) GetUserByID(id uint) (error, *user_model.User) {
	return userService.iUser.GetUserByID(id)
}
