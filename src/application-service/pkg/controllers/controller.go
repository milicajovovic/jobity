package controllers

import (
	"application-service/pkg/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", GetAll)
	app.Get("/application/:id", GetById)
	app.Post("/apply/:adId/:employeeId", Apply)
	app.Get("/accepted/:employeeId", GetAccepted)
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

func Apply(c *fiber.Ctx) error {
	paramId := c.Params("adId")
	adId, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ad ID")
	}

	paramId = c.Params("employeeId")
	employeeId, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid employee ID")
	}

	ad, err := services.Apply(adId, employeeId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ad)
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
