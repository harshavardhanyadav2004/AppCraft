package infrastructure

import (
	"fmt"
	"log"

	"appcraft_web/app_models"
	"appcraft_web/metrics_models"
	"appcraft_web/signup_models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "postgres", "appcraft", "5432")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&app_models.App{}, &metrics_models.Metrics{}, &metrics_models.MetricTypes{},
		&metrics_models.DailyMetrics{}, &metrics_models.WeeklyMetrics{}, &metrics_models.MonthlyMetrics{},
		&metrics_models.YearlyMetrics{}, &signup_models.Signup{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}