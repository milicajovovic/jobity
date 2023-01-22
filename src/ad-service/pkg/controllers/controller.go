package controllers

import (
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
		requiredSkills = append(requiredSkills, ad.RequierdSkills...)
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
