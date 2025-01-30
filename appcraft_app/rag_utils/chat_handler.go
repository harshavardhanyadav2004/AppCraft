package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"your-project/rag_utils"
)

type ChatRequest struct {
	Query string `json:"query"`
}

func ChatHandler(c *fiber.Ctx) error {
	var chatRequest ChatRequest
	err := c.BodyParser(&chatRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	answer, err := rag_utils.ExtractAnswer(chatRequest.Query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString(answer)
}