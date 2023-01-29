package routes

import (
	"application-service/pkg/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	"github.com/gofiber/fiber/v2"
)

const applicationPrefix = "/applications"
const applicationUrl = "http://localhost:3002"

func SetupApplicationRoutes(app *fiber.App, auth *fibercasbin.CasbinMiddleware) {
	// Get all applications
	app.Get(applicationPrefix+"/", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		response, err := http.Get(applicationUrl)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var applications []models.Application
		err = json.NewDecoder(response.Body).Decode(&applications)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(applications)
	})

	// Get application by id
	app.Get(applicationPrefix+"/application/:id", auth.RequiresRoles([]string{"employer"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Get(applicationUrl + "/application/" + paramId)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var application models.Application
		err = json.NewDecoder(response.Body).Decode(&application)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(application)
	})

	// Apply to an ad - create application
	app.Post(applicationPrefix+"/apply/:adId/:employeeId", auth.RequiresRoles([]string{"employee"}), func(c *fiber.Ctx) error {
		adId := c.Params("adId")
		employeeId := c.Params("employeeId")
		response, err := http.Post(applicationUrl+"/apply/"+adId+"/"+employeeId, "application/json", bytes.NewBuffer(c.Body()))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var application models.Application
		err = json.NewDecoder(response.Body).Decode(&application)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(application)
	})

	// Get accepted applications for employee
	app.Get(applicationPrefix+"/accepted/:id", auth.RequiresRoles([]string{"employee"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Get(applicationUrl + "/accepted/" + paramId)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var applications []models.Application
		err = json.NewDecoder(response.Body).Decode(&applications)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(applications)
	})
}
