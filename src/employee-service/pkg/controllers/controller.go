package controllers

import (
	"employee-service/pkg/models"
	"employee-service/pkg/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", GetAll)
	app.Get("/employee/:id", GetById)
	app.Post("/register/form", RegisterForm)
	app.Post("/register/pdf", RegisterPdf)
	app.Post("/update/form", UpdateForm)
	app.Post("/update/pdf", UpdatePdf)
	app.Post("/login", Login)
	app.Post("/block/:id", Block)
	app.Post("/delete/:id", Delete)
}

func GetAll(c *fiber.Ctx) error {
	employees, err := services.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employees)
}

func GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	employee, err := services.GetById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func RegisterForm(c *fiber.Ctx) error {
	var newEmployee models.Employee
	if err := c.BodyParser(&newEmployee); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	employee, err := services.RegisterForm(newEmployee)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func RegisterPdf(c *fiber.Ctx) error {
	var dto models.RegisterDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	employee, err := services.RegisterPdf(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func UpdateForm(c *fiber.Ctx) error {
	var updatedEmployee models.Employee
	if err := c.BodyParser(&updatedEmployee); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	employee, err := services.UpdateForm(updatedEmployee)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func UpdatePdf(c *fiber.Ctx) error {
	var dto models.RegisterDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	employee, err := services.UpdatePdf(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func Login(c *fiber.Ctx) error {
	var dto models.LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	employee, err := services.Login(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func Block(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	employee, err := services.Block(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}

func Delete(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	employee, err := services.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(employee)
}
