package mobile_handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/mobile/apps", CreateMobileApp)
	app.Get("/mobile/apps/:appID", GetMobileApp)
	app.Put("/mobile/apps/:appID", UpdateMobileApp)
	app.Delete("/mobile/apps/:appID", DeleteMobileApp)
}