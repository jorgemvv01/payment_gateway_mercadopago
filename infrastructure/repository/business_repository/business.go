package business_repository

import (
	"errors"
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/business_model"
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

func (businessRepository Repository) GetBusinessByID(id uint) (error, *business_model.Business) {
	var business *business_model.Business
	if err := businessRepository.db.Find(&business, id).Error; err != nil {
		return err, nil
	}
	if business.ID == 0 {
		return errors.New("business with ID: " + strconv.Itoa(int(id)) + " not found"), nil
	}
	return nil, business
}

func (businessRepository Repository) GetAllBusiness() (error, *[]business_model.Business) {
	var business *[]business_model.Business
	if err := businessRepository.db.Find(&business).Error; err != nil {
		return err, nil
	}
	return nil, business
}
