package user_model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Email    string `gorm:"not null"`
}
