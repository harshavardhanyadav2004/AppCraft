package login_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/login_models"
)

func DeleteUser(id string) error {
	db := infrastructure.GetDB()
	var user login_models.User
	return db.Where("id = ?", id).Delete(&user).Error
}