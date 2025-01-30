package metrics_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/metrics_models"
)

func CreateMetric(metric metrics_models.Metrics) error {
	db := infrastructure.GetDB()
	return db.Create(&metric).Error
}