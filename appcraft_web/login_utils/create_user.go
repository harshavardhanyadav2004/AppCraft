package login_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/login_models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user login_models.User) error {
	db := infrastructure.GetDB()
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return db.Create(&user).Error
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}