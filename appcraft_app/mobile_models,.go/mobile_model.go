package mobile_models

import (
	"gorm.io/gorm"
)

type MobileApp struct {
	gorm.Model
	AppID      string `json:"app_id" gorm:"unique"`
	AppName    string `json:"app_name"`
	AppType    string `json:"app_type"`
	Component  string `json:"component"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}