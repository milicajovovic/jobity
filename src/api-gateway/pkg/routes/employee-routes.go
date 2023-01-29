package routes

import (
	apimodels "api-gateway/pkg/models"
	"api-gateway/pkg/utils"
	"bytes"
	"employee-service/pkg/models"
	"encoding/json"
	"io"
	"net/http"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

const employeePrefix = "/employees"
const employeeUrl = "http://localhost:3003"

func SetupEmployeeRoutes(app *fiber.App, auth *fibercasbin.CasbinMiddleware, enforcer *casbin.Enforcer) {
	// Get all employees
	app.Get(employeePrefix+"/", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		response, err := http.Get(employeeUrl)

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
	})

	// Get employee by id
	app.Get(employeePrefix+"/employee/:id", auth.RequiresRoles([]string{"employee"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Get(employeeUrl + "/employee/" + paramId)

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
	})

	// Register employee through form
	app.Post(employeePrefix+"/register/form", func(c *fiber.Ctx) error {
		response, err := http.Post(employeeUrl+"/register/form", "application/json", bytes.NewBuffer(c.Body()))

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

		enforcer.AddGroupingPolicy(employee.Email, "employee")
		token, _ := utils.GenerateToken(employee.ID, employee.Email)
		dto := apimodels.LoggedInDto{
			Jwt:    token,
			UserId: employee.ID,
		}

		return c.Status(response.StatusCode).JSON(dto)
	})

	// Register employee through pdf
	app.Post(employeePrefix+"/register/pdf", func(c *fiber.Ctx) error {
		response, err := http.Post(employeeUrl+"/register/pdf", "application/json", bytes.NewBuffer(c.Body()))

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

		enforcer.AddGroupingPolicy(employee.Email, "employee")

		token, _ := utils.GenerateToken(employee.ID, employee.Email)
		dto := apimodels.LoggedInDto{
			Jwt:    token,
			UserId: employee.ID,
		}

		return c.Status(response.StatusCode).JSON(dto)
	})

	// Update employee throught form
	app.Post(employeePrefix+"/update/form", auth.RequiresRoles([]string{"employee"}), func(c *fiber.Ctx) error {
		response, err := http.Post(employeeUrl+"/update/form", "application/json", bytes.NewBuffer(c.Body()))

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
	})

	// Update employee throught pdf
	app.Post(employeePrefix+"/update/pdf", auth.RequiresRoles([]string{"employee"}), func(c *fiber.Ctx) error {
		response, err := http.Post(employeeUrl+"/update/pdf", "application/json", bytes.NewBuffer(c.Body()))

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
	})

	// Login as an employee
	app.Post(employeePrefix+"/login", func(c *fiber.Ctx) error {
		response, err := http.Post(employeeUrl+"/login", "application/json", bytes.NewBuffer(c.Body()))

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

		token, _ := utils.GenerateToken(employee.ID, employee.Email)
		dto := apimodels.LoggedInDto{
			Jwt:    token,
			UserId: employee.ID,
		}

		return c.Status(response.StatusCode).JSON(dto)
	})
}
