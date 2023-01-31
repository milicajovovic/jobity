package controllers

import (
	"employer-service/pkg/models"
	"employer-service/pkg/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", GetAll)
	app.Get("/employer/:id", GetById)
	app.Post("/register", Register)
	app.Post("/login", Login)
	app.Post("/update", Update)
	app.Post("/delete/:id", Delete)
}

func GetAll(c *fiber.Ctx) error {
	employers, err := services.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employers)
}

func GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	employer, err := services.GetById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employer)
}

func Register(c *fiber.Ctx) error {
	var newEmployer models.Employer
	if err := c.BodyParser(&newEmployer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	employer, err := services.Register(newEmployer)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employer)
}

func Login(c *fiber.Ctx) error {
	var dto models.LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	employer, err := services.Login(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employer)
}

func Update(c *fiber.Ctx) error {
	var updatedEmployer models.Employer
	if err := c.BodyParser(&updatedEmployer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	employer, err := services.Update(updatedEmployer)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employer)
}

func Delete(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	employer, err := services.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employer)
}
