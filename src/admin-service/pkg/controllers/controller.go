package controllers

import (
	"admin-service/pkg/models"
	"admin-service/pkg/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", GetAll)
	app.Get("/admin/:id", GetById)
	app.Post("/login", Login)
}

func GetAll(c *fiber.Ctx) error {
	admins, err := services.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(admins)
}

func GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	admin, err := services.GetById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(admin)
}

func Login(c *fiber.Ctx) error {
	var dto models.LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	admin, err := services.Login(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(admin)
}
