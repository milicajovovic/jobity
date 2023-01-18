package main

import (
	"ad-service/pkg/config"
	"ad-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()

	app := fiber.New()
	controllers.SetupRoutes(app)
	app.Listen(":3001")
}
