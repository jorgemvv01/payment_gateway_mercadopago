package business_model

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	MpToken     string
}

type BusinessResponse struct {
	ID          uint
	Name        string
	Description string
}

func NewBusinessResponse(business Business) BusinessResponse {
	return BusinessResponse{
		ID:          business.ID,
		Name:        business.Name,
		Description: business.Description,
	}
}
