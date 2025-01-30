package routes
import (
	"github.com/gofiber/fiber/v2"
	"your-project/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/chat", handlers.ChatHandler)
}