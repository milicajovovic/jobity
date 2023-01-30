package routes

import (
	"ad-service/pkg/models"
	"encoding/json"
	"io"
	"net/http"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	"github.com/gofiber/fiber/v2"
)

const adPrefix = "/ads"
const adUrl = "http://localhost:3001"

func SetupAdRoutes(app *fiber.App, auth *fibercasbin.CasbinMiddleware) {
	// Get all ads
	app.Get(adPrefix+"/", func(c *fiber.Ctx) error {
		response, err := http.Get(adUrl)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var ads []models.Ad
		err = json.NewDecoder(response.Body).Decode(&ads)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(ads)
	})

	// Get ad by id
	app.Get(adPrefix+"/ad/:id", auth.RequiresRoles([]string{"employee"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Get(adUrl + "/ad/" + paramId)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var ad models.Ad
		err = json.NewDecoder(response.Body).Decode(&ad)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(ad)
	})

	// Search ads
	app.Get(adPrefix+"/search/:name/:description", func(c *fiber.Ctx) error {
		name := c.Params("name")
		description := c.Params("description")
		response, err := http.Get(adUrl + "/search/" + name + "/" + description)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var ads []models.Ad
		err = json.NewDecoder(response.Body).Decode(&ads)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(ads)
	})

	// Get all job types
	app.Get(adPrefix+"/jobTypes", func(c *fiber.Ctx) error {
		response, err := http.Get(adUrl + "/jobTypes")

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var jobTypes []string
		err = json.NewDecoder(response.Body).Decode(&jobTypes)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(jobTypes)
	})

	// Get all skills
	app.Get(adPrefix+"/requiredSkills", func(c *fiber.Ctx) error {
		response, err := http.Get(adUrl + "/requiredSkills")

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var requiredSkills []string
		err = json.NewDecoder(response.Body).Decode(&requiredSkills)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(requiredSkills)
	})

	// Delete ad
	app.Post(adPrefix+"/delete/:id", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Post(adUrl+"/delete/"+paramId, "application/json", nil)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var ad models.Ad
		err = json.NewDecoder(response.Body).Decode(&ad)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(ad)
	})
}
