package metric_handlers

import (
	"appcraft_web/metrics_models"
	"appcraft_web/metrics_utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func CreateMetric(c *fiber.Ctx) error {
	var metric metrics_models.Metrics
	err := json.Unmarshal(c.Body(), &metric)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = metrics_utils.CreateMetric(metric)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}