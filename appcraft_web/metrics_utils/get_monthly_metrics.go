package metrics_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/metrics_models"

	"time"
)

func GetMonthlyMetrics(appID uint, component string) ([]metrics_models.MonthlyMetrics, error) {
	db := infrastructure.GetDB()
	var monthlyMetrics []metrics_models.MonthlyMetrics
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	err := db.Model(&metrics_models.Metrics{}).
		Where("app_id = ? AND component = ? AND date BETWEEN ? AND ?", appID, component, startOfMonth.Format("2006-01-02"), endOfMonth.Format("2006-01-02")).
		Find(&monthlyMetrics).Error
	return monthlyMetrics, err
}