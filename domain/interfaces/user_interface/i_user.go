package user_interface

import "payment_gateway_mercadopago/domain/models/user_model"

type IUser interface {
	GetUserByID(id uint) (error, *user_model.User)
}
