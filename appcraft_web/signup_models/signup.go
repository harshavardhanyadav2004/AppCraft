package signup_models

import (
	"gorm.io/gorm"
)

type Signup struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"unique;not null"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null;email"`
	PhoneNo   int    `json:"phone_no" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	ConfirmPassword string `json:"confirm_password" gorm:"not null"`
}