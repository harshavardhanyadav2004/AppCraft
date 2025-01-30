package metric_handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/metrics", CreateMetric)
	app.Get("/metrics/daily/:appID/:component", GetDailyMetrics)
	app.Get("/metrics/weekly/:appID/:component", GetWeeklyMetrics)
	app.Get("/metrics/monthly/:appID/:component", GetMonthlyMetrics)
	app.Get("/metrics/yearly/:appID/:component", GetYearlyMetrics)
}