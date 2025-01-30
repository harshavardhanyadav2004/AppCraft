package metrics_models

import (
	"gorm.io/gorm"
)

type Metrics struct {
	gorm.Model
	AppID      uint   `json:"app_id"`
	Component  string `json:"component"`
	Date       string `json:"date"`
	Hours      int    `json:"hours"`
	MetricType string `json:"metric_type"`
	Value      int    `json:"value"`
}

type MetricTypes struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	MetricType  string `json:"metric_type"`
	AppID       uint   `json:"app_id"`
	Component   string `json:"component"`
}

type DailyMetrics struct {
	Date       string `json:"date"`
	Hours      int    `json:"hours"`
	MetricType string `json:"metric_type"`
	Value      int    `json:"value"`
}

type WeeklyMetrics struct {
	Week       string `json:"week"`
	Hours      int    `json:"hours"`
	MetricType string `json:"metric_type"`
	Value      int    `json:"value"`
}

type MonthlyMetrics struct {
	Month      string `json:"month"`
	Hours      int    `json:"hours"`
	MetricType string `json:"metric_type"`
	Value      int    `json:"value"`
}

type YearlyMetrics struct {
	Year       string `json:"year"`
	Hours      int    `json:"hours"`
	MetricType string `json:"metric_type"`
	Value      int    `json:"value"`
}