package main

import (
	"appcraft_web/infrastructure"
	"appcraft_web/login_handlers"
	"appcraft_web/signup_handlers"
	"appcraft_web/metric_handlers"
	"appcraft_web/rag_handlers"
	"appcraft_web/mobile_handlers"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	infrastructure.InitDB()
	defer infrastructure.CloseDB()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Static("/", "./public")

	app.Group("/api")
	app.Group("/api/metrics")
	metric_handlers.RegisterRoutes(app)
	app.Group("/api")
	login_handlers.RegisterRoutes(app)
	signup_handlers.RegisterRoutes(app)
	rag_handlers.RegisterRoutes(app)
	mobile_handlers.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}