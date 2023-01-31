package controllers

import (
	"ad-service/pkg/models"
	"ad-service/pkg/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", GetAll)
	app.Get("/ad/:id", GetById)
	app.Get("/search/:name/:description", Search)
	app.Get("/jobTypes", GetJobTypes)
	app.Get("/requiredSkills", GetRequiredSkills)
	app.Post("/update", Update)
	app.Post("/delete/:id", Delete)
	app.Get("/employer/:id", GetByEmployerId)
	app.Post("/create", Create)
}

func GetAll(c *fiber.Ctx) error {
	ads, err := services.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ads)
}

func GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	ad, err := services.GetById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ad)
}

func Search(c *fiber.Ctx) error {
	name := c.Params("name")
	description := c.Params("description")

	ads, err := services.Search(name, description)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ads)
}

func GetJobTypes(c *fiber.Ctx) error {
	ads, err := services.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var jobTypes []string
	for _, ad := range ads {
		jobTypes = append(jobTypes, ad.JobType...)
	}
	jobTypes = removeDuplicates(jobTypes)

	return c.Status(fiber.StatusOK).JSON(jobTypes)
}

func GetRequiredSkills(c *fiber.Ctx) error {
	ads, err := services.GetAll()

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var requiredSkills []string
	for _, ad := range ads {
		requiredSkills = append(requiredSkills, ad.RequiredSkills...)
	}
	requiredSkills = removeDuplicates(requiredSkills)

	return c.Status(fiber.StatusOK).JSON(requiredSkills)
}

func removeDuplicates(slice []string) []string {
	allStrings := make(map[string]bool)
	unique := []string{}
	for _, item := range slice {
		if _, value := allStrings[item]; !value {
			allStrings[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

func Update(c *fiber.Ctx) error {
	var updatedAd models.Ad
	if err := c.BodyParser(&updatedAd); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	ad, err := services.Update(updatedAd)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ad)
}

func Delete(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	ad, err := services.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ad)
}

func GetByEmployerId(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	ads, err := services.GetByEmployerId(id)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ads)
}

func Create(c *fiber.Ctx) error {
	var newAd models.Ad
	if err := c.BodyParser(&newAd); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	ad, err := services.Create(newAd)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(ad)
}
