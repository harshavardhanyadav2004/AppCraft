package metric_handlers

import (
	"appcraft_web/metrics_utils"

	"github.com/gofiber/fiber/v2"
)

func GetMonthlyMetrics(c *fiber.Ctx) error {
	appID, err := c.ParamsInt("appID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	component := c.Params("component")

	monthlyMetrics, err := metrics_utils.GetMonthlyMetrics(uint(appID), component)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(monthlyMetrics)
}