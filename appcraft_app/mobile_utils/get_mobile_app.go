package mobile_utils

import (
	"appcraft_app/infrastructure"
	"appcraft_app/mobile_models"
	"gorm.io/gorm"
)

func GetMobileApp(appID string) (mobile_models.MobileApp, error) {
	db := infrastructure.GetDB()
	var app mobile_models.MobileApp
	err := db.Where("app_id = ?", appID).First(&app).Error
	return app, err
}