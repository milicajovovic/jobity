package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"review-service/pkg/models"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	"github.com/gofiber/fiber/v2"
)

const reviewPrefix = "/reviews"
const reviewUrl = "http://localhost:3005"

func SetupReviewRoutes(app *fiber.App, auth *fibercasbin.CasbinMiddleware) {
	// Get all reviews
	app.Get(reviewPrefix+"/", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		response, err := http.Get(reviewUrl)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var reviews []models.Review
		err = json.NewDecoder(response.Body).Decode(&reviews)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(reviews)
	})

	// Get employer's reviews
	app.Get(reviewPrefix+"/employer/:id", auth.RequiresRoles([]string{"employer"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Get(reviewUrl + "/employer/" + paramId)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var reviews []models.Review
		err = json.NewDecoder(response.Body).Decode(&reviews)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(reviews)
	})

	// Create review
	app.Post(reviewPrefix+"/create", auth.RequiresRoles([]string{"employee"}), func(c *fiber.Ctx) error {
		response, err := http.Post(reviewUrl+"/create", "application/json", bytes.NewBuffer(c.Body()))

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var review models.Review
		err = json.NewDecoder(response.Body).Decode(&review)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(review)
	})

	// Mark review as appropriate
	app.Post(reviewPrefix+"/appropriate/:id", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Post(reviewUrl+"/appropriate/"+paramId, "application/json", nil)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var review models.Review
		err = json.NewDecoder(response.Body).Decode(&review)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(review)
	})

	// Mark review as inappropriate
	app.Post(reviewPrefix+"/inappropriate/:id", auth.RequiresRoles([]string{"employer"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Post(reviewUrl+"/inappropriate/"+paramId, "application/json", nil)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var review models.Review
		err = json.NewDecoder(response.Body).Decode(&review)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(review)
	})

	// Delete review
	app.Post(reviewPrefix+"/delete/:id", auth.RequiresRoles([]string{"admin"}), func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		response, err := http.Post(reviewUrl+"/delete/"+paramId, "application/json", nil)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			body, _ := io.ReadAll(response.Body)
			return fiber.NewError(response.StatusCode, string(body))
		}

		var review models.Review
		err = json.NewDecoder(response.Body).Decode(&review)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.Status(response.StatusCode).JSON(review)
	})
}
