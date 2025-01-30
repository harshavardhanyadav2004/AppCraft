package metrics_utils

import (
	"appcraft_web/infrastructure"
	"appcraft_web/metrics_models"
	"time"
)

func GetWeeklyMetrics(appID uint, component string) ([]metrics_models.WeeklyMetrics, error) {
	db := infrastructure.GetDB()
	var weeklyMetrics []metrics_models.WeeklyMetrics
	now := time.Now()
	startOfWeek := now.Add(-time.Duration(now.Weekday()) * time.Duration(time.Now().Day()))
	endOfWeek := startOfWeek.Add(7 * time.Duration(time.Now().Day()))
	err := db.Model(&metrics_models.Metrics{}).
		Where("app_id = ? AND component = ? AND date BETWEEN ? AND ?", appID, component, startOfWeek.Format("2006-01-02"), endOfWeek.Format("2006-01-02")).
		Find(&weeklyMetrics).Error
	return weeklyMetrics, err
}