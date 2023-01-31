package controllers

import (
	"application-service/pkg/models"
	"application-service/pkg/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", GetAll)
	app.Get("/application/:id", GetById)
	app.Get("/employer/:id", GetByEmployerId)
	app.Post("/apply", Apply)
	app.Get("/accepted/:employeeId", GetAccepted)
	app.Post("/update", Update)
	app.Get("/interviews/:id", GetInterviews)
}

func GetAll(c *fiber.Ctx) error {
	applications, err := services.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(applications)
}

func GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	application, err := services.GetById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(application)
}

func GetByEmployerId(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	applications, err := services.GetByEmployerId(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(applications)
}

func Apply(c *fiber.Ctx) error {
	var newApplication models.Application
	if err := c.BodyParser(&newApplication); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	application, err := services.Apply(newApplication)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(application)
}

func GetAccepted(c *fiber.Ctx) error {
	paramId := c.Params("employeeId")
	employeeId, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	applications, err := services.GetAccepted(employeeId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(applications)
}

func Update(c *fiber.Ctx) error {
	var updatedApplication models.Application
	if err := c.BodyParser(&updatedApplication); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	application, err := services.Update(updatedApplication)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(application)
}

func GetInterviews(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	applications, err := services.GetInterviews(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(applications)
}
