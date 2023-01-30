package main

import (
	"admin-service/pkg/config"
	"admin-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()

	app := fiber.New()
	controllers.SetupRoutes(app)
	app.Listen(":3000")
}
