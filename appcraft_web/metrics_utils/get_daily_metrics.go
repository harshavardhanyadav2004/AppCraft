package metrics_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/metrics_models"
	"time"
)

func GetDailyMetrics(appID uint, component string) ([]metrics_models.DailyMetrics, error) {
	db := infrastructure.GetDB()
	var dailyMetrics []metrics_models.DailyMetrics
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	err := db.Model(&metrics_models.Metrics{}).
		Where("app_id = ? AND component = ? AND date BETWEEN ? AND ?", appID, component, startOfDay.Format("2006-01-02"), endOfDay.Format("2006-01-02")).
		Find(&dailyMetrics).Error
	return dailyMetrics, err
}