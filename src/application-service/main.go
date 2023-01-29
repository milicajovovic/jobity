package main

import (
	"application-service/pkg/config"
	"application-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()

	app := fiber.New()
	controllers.SetupRoutes(app)
	app.Listen(":3002")
}
