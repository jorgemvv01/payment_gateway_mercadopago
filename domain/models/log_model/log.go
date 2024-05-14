package log_model

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	LogType     string
	Status      string
	Information string
	Details     string
	Module      string
}
