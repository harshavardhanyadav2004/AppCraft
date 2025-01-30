package metrics_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/metrics_models"
	"time"
)

func GetYearlyMetrics(appID uint, component string) ([]metrics_models.YearlyMetrics, error) {
	db := infrastructure.GetDB()
	var yearlyMetrics []metrics_models.YearlyMetrics
	now := time.Now()
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	endOfYear := startOfYear.AddDate(1, 0, -1)
	err := db.Model(&metrics_models.Metrics{}).
		Where("app_id = ? AND component = ? AND date BETWEEN ? AND ?", appID, component, startOfYear.Format("2006-01-02"), endOfYear.Format("2006-01-02")).
		Find(&yearlyMetrics).Error
	return yearlyMetrics, err
}