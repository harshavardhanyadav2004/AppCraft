package signup_handlers

import (
	"appcraft_web/signup_models"
	"appcraft_web/signup_utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func CreateSignup(c *fiber.Ctx) error {
	var signup signup_models.Signup
	err := json.Unmarshal(c.Body(), &signup)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = signup_utils.ValidateSignup(signup)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = signup_utils.CreateSignup(signup)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}