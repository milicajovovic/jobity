package routes

import (
	"admin-service/pkg/models"
	apimodels "api-gateway/pkg/models"
	"api-gateway/pkg/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	"github.com/gofiber/fiber/v2"
)

const adminPrefix = "/admins"
const adminUrl = "http://localhost:3000"

func SetupAdminRoutes(app *fiber.App, auth *fibercasbin.CasbinMiddleware) {
	// Get all admins
	app.Get(adminPrefix+"/", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		response, err := http.Get(adminUrl)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var admins []models.Admin
		err = json.NewDecoder(response.Body).Decode(&admins)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(admins)
	})

	// Get admin by id
	app.Get(adminPrefix+"/admin/:id", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Get(adminUrl + "/admin/" + paramId)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var admin models.Admin
		err = json.NewDecoder(response.Body).Decode(&admin)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(admin)
	})

	// Login as an admin
	app.Post(adminPrefix+"/login", func(c *fiber.Ctx) error {
		response, err := http.Post(adminUrl+"/login", "application/json", bytes.NewBuffer(c.Body()))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var admin models.Admin
		err = json.NewDecoder(response.Body).Decode(&admin)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		token, _ := utils.GenerateToken(admin.ID, admin.Email)
		dto := apimodels.LoggedInDto{
			Jwt:    token,
			UserId: admin.ID,
		}

		return c.Status(response.StatusCode).JSON(dto)
	})
}
