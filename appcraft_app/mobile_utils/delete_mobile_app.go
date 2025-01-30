package mobile_utils

import (
	"appcraft_app/infrastructure"
	"appcraft_app/mobile_models"
	"gorm.io/gorm"
)

func DeleteMobileApp(appID string) error {
	db := infrastructure.GetDB()
	var app mobile_models.MobileApp
	return db.Where("app_id = ?", appID).Delete(&app).Error
}