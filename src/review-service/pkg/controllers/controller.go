package controllers

import (
	"review-service/pkg/models"
	"review-service/pkg/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", GetAll)
	app.Get("/employer/:id", GetByEmployerId)
	app.Post("/create", Create)
}

func GetAll(c *fiber.Ctx) error {
	reviews, err := services.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(reviews)
}

func GetByEmployerId(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	reviews, err := services.GetByEmployerId(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(reviews)
}

func Create(c *fiber.Ctx) error {
	var newReview models.Review
	if err := c.BodyParser(&newReview); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	ad, err := services.Create(newReview)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ad)
}
