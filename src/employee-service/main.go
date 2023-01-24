package main

import (
	"employee-service/pkg/config"
	"employee-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3006",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	controllers.SetupRoutes(app)
	app.Listen(":3002")
}
