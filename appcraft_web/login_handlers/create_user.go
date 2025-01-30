package login_handlers

import (
	"appcraft_web/login_models"
	"appcraft_web/login_utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user login_models.User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = login_utils.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}