package user_repository

import (
	"errors"
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/user_model"
	"strconv"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (userRepository Repository) GetUserByID(id uint) (error, *user_model.User) {
	var user *user_model.User
	if err := userRepository.db.Find(&user, id).Error; err != nil {
		return err, nil
	}
	if user.ID == 0 {
		return errors.New("user with ID: " + strconv.Itoa(int(id)) + " not found"), nil
	}
	return nil, user
}
