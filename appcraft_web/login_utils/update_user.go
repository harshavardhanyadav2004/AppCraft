package login_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/login_models"
)

func UpdateUser(user login_models.User) error {
	db := infrastructure.GetDB()
	return db.Model(&login_models.User{}).Updates(user).Error
}