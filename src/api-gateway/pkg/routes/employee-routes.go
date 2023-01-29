package routes

import (
	"api-gateway/pkg/utils"
	"bytes"
	"employee-service/pkg/models"
	"encoding/json"
	"io"
	"net/http"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	"github.com/gofiber/fiber/v2"
)

const apiPrefix = "/employees"
const url = "http://localhost:3003"

func SetupEmployeeRoutes(app *fiber.App, auth *fibercasbin.CasbinMiddleware) {
	app.Get(apiPrefix+"/", auth.RequiresRoles([]string{"admin"}), GetAll)
	app.Get(apiPrefix+"/:id", GetById)
	app.Post(apiPrefix+"/register/form", RegisterForm)
	app.Post(apiPrefix+"/register/pdf", RegisterPdf)
	app.Post(apiPrefix+"/update", auth.RequiresRoles([]string{"employee"}), Update)
	app.Post(apiPrefix+"/login", Login)
}

func GetAll(c *fiber.Ctx) error {
	response, err := http.Get(url)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	defer response.Body.Close()

	if response.Status != "200 OK" {
		body, _ := io.ReadAll(response.Body)
		return fiber.NewError(response.StatusCode, string(body))
	}

	var employees []models.Employee
	err = json.NewDecoder(response.Body).Decode(&employees)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(response.StatusCode).JSON(employees)
}

func GetById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	response, err := http.Get(url + "/employee/" + paramId)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	defer response.Body.Close()

	if response.Status != "200 OK" {
		body, _ := io.ReadAll(response.Body)
		return fiber.NewError(response.StatusCode, string(body))
	}

	var employee models.Employee
	err = json.NewDecoder(response.Body).Decode(&employee)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(response.StatusCode).JSON(employee)
}

func RegisterForm(c *fiber.Ctx) error {
	response, err := http.Post(url+"/register/form", "application/json", bytes.NewBuffer(c.Body()))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	defer response.Body.Close()

	if response.Status != "200 OK" {
		body, _ := io.ReadAll(response.Body)
		return fiber.NewError(response.StatusCode, string(body))
	}

	var employee models.Employee
	err = json.NewDecoder(response.Body).Decode(&employee)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(response.StatusCode).JSON(employee)
}

func RegisterPdf(c *fiber.Ctx) error {
	response, err := http.Post(url+"/register/pdf", "application/json", bytes.NewBuffer(c.Body()))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	defer response.Body.Close()

	if response.Status != "200 OK" {
		body, _ := io.ReadAll(response.Body)
		return fiber.NewError(response.StatusCode, string(body))
	}

	var employee models.Employee
	err = json.NewDecoder(response.Body).Decode(&employee)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(response.StatusCode).JSON(employee)
}

func Update(c *fiber.Ctx) error {
	response, err := http.Post(url+"/update", "application/json", bytes.NewBuffer(c.Body()))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	defer response.Body.Close()

	if response.Status != "200 OK" {
		body, _ := io.ReadAll(response.Body)
		return fiber.NewError(response.StatusCode, string(body))
	}

	var employee models.Employee
	err = json.NewDecoder(response.Body).Decode(&employee)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(response.StatusCode).JSON(employee)
}

func Login(c *fiber.Ctx) error {
	response, err := http.Post(url+"/login", "application/json", bytes.NewBuffer(c.Body()))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	defer response.Body.Close()

	if response.Status != "200 OK" {
		body, _ := io.ReadAll(response.Body)
		return fiber.NewError(response.StatusCode, string(body))
	}

	var employee models.Employee
	err = json.NewDecoder(response.Body).Decode(&employee)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	token, _ := utils.GenerateToken(employee.ID, employee.Email, "employee")
	return c.Status(response.StatusCode).JSON(token)
}
