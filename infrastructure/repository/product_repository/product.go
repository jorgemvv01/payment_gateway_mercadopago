package product_repository

import (
	"gorm.io/gorm"
	"payment_gateway_mercadopago/domain/models/product_model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (productRepository Repository) GetProductsByIDs(ids []uint) (error, *[]product_model.Product) {
	var products *[]product_model.Product
	if err := productRepository.db.Where("id IN (?)", ids).Find(&products).Error; err != nil {
		return err, nil
	}
	return nil, products
}

func (productRepository Repository) GetAllProductsByBusiness(businessID uint) (error, *[]product_model.Product) {
	var products *[]product_model.Product
	if err := productRepository.db.Where("business_id = ?", businessID).Find(&products).Error; err != nil {
		return err, nil
	}
	return nil, products
}

func (productRepository Repository) GetPromotionalProductsByBusiness(businessID uint) (error, *[]product_model.Product) {
	var products *[]product_model.Product
	if err := productRepository.db.Where("business_id = ? and discount > ?", businessID, 0).Order("discount desc").Limit(10).Find(&products).Error; err != nil {
		return err, nil
	}
	return nil, products
}
