package main

import (
	"employee-service/pkg/config"
	"employee-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()

	app := fiber.New()
	controllers.SetupRoutes(app)
	app.Listen(":3002")
}
