package mobile_handlers

import (
	"appcraft_app/mobile_utils"
	"github.com/gofiber/fiber/v2"
)

func DeleteMobileApp(c *fiber.Ctx) error {
	appID := c.Params("appID")
	err := mobile_utils.DeleteMobileApp(appID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}