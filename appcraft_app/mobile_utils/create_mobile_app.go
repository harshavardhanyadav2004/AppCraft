package mobile_utils

import (
	"appcraft_app/infrastructure"
	"appcraft_app/mobile_models"
	"gorm.io/gorm"
)

func CreateMobileApp(app mobile_models.MobileApp) error {
	db := infrastructure.GetDB()
	return db.Create(&app).Error
}