package routes

import (
	apimodels "api-gateway/pkg/models"
	"api-gateway/pkg/utils"
	"bytes"
	"employer-service/pkg/models"
	"encoding/json"
	"io"
	"net/http"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

const employerPrefix = "/employers"
const employerUrl = "http://localhost:3004"

func SetupEmployerRoutes(app *fiber.App, auth *fibercasbin.CasbinMiddleware, enforcer *casbin.Enforcer) {
	// Get all employers
	app.Get(employerPrefix+"/", func(c *fiber.Ctx) error {
		response, err := http.Get(employerUrl)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var employers []models.Employer
		err = json.NewDecoder(response.Body).Decode(&employers)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(employers)
	})

	// Get employer by id
	app.Get(employerPrefix+"/employer/:id", func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Get(employerUrl + "/employer/" + paramId)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var employer models.Employer
		err = json.NewDecoder(response.Body).Decode(&employer)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(employer)
	})

	// Register employer
	app.Post(employerPrefix+"/register", func(c *fiber.Ctx) error {
		response, err := http.Post(employerUrl+"/register", "application/json", bytes.NewBuffer(c.Body()))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var employer models.Employer
		err = json.NewDecoder(response.Body).Decode(&employer)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		enforcer.AddGroupingPolicy(employer.Email, "employer")

		token, _ := utils.GenerateToken(employer.ID, employer.Email)
		dto := apimodels.LoggedInDto{
			Jwt:    token,
			UserId: employer.ID,
		}
		return c.Status(response.StatusCode).JSON(dto)
	})

	// Login as an employer
	app.Post(employerPrefix+"/login", func(c *fiber.Ctx) error {
		response, err := http.Post(employerUrl+"/login", "application/json", bytes.NewBuffer(c.Body()))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var employer models.Employer
		err = json.NewDecoder(response.Body).Decode(&employer)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		token, _ := utils.GenerateToken(employer.ID, employer.Email)
		dto := apimodels.LoggedInDto{
			Jwt:    token,
			UserId: employer.ID,
		}

		return c.Status(response.StatusCode).JSON(dto)
	})

	// Update employer
	app.Post(employerPrefix+"/update", auth.RequiresRoles([]string{"employer"}), func(c *fiber.Ctx) error {
		response, err := http.Post(employerUrl+"/update", "application/json", bytes.NewBuffer(c.Body()))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var employer models.Employer
		err = json.NewDecoder(response.Body).Decode(&employer)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(employer)
	})

	// Delete employer
	app.Post(employerPrefix+"/delete/:id", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Post(employerUrl+"/delete/"+paramId, "application/json", nil)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var employer models.Employer
		err = json.NewDecoder(response.Body).Decode(&employer)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(employer)
	})
}
