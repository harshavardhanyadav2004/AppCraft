package mobile_handlers

import (
	"appcraft_app/mobile_models"
	"appcraft_app/mobile_utils"
	"github.com/gofiber/fiber/v2"
)

func GetMobileApp(c *fiber.Ctx) error {
	appID := c.Params("appID")
	app, err := mobile_utils.GetMobileApp(appID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(app)
}