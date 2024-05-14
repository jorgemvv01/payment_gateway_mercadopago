package log_repository

import (
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/log_model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (logRepository Repository) CreateLog(log *log_model.Log) error {
	return logRepository.db.Create(&log).Error
}
