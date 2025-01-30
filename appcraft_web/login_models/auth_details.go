package login_models

import (
	"gorm.io/gorm"
)

type AuthDetails struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
}