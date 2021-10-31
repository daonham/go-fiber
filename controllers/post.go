package controllers

import (
	"github.com/daonham/go-fiber/forms"
	"github.com/daonham/go-fiber/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func CreatePost(c *fiber.Ctx) error {

	create := &forms.PostForm{}

	if err := c.BodyParser(create); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate = validator.New()

	err := validate.Struct(create)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	id, message, err := models.CreatePost(create)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": message,
		"id":      id,
	})
}
