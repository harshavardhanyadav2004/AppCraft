package signup_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/signup_models"
	"golang.org/x/crypto/bcrypt"
)

func CreateSignup(signup signup_models.Signup) error {
	db := infrastructure.GetDB()
	hashedPassword, err := hashPassword(signup.Password)
	if err != nil {
		return err
	}
	signup.Password = hashedPassword
	return db.Create(&signup).Error
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}