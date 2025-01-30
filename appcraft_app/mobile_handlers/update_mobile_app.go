package mobile_handlers

import (
	"appcraft_app/mobile_models"
	"appcraft_app/mobile_utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func UpdateMobileApp(c *fiber.Ctx) error {
	var app mobile_models.MobileApp
	err := json.Unmarshal(c.Body(), &app)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = mobile_utils.UpdateMobileApp(app)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}