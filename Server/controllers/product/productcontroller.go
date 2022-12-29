package product

import (
	"Server/Api/database"
	"Server/Api/models"

	"github.com/gofiber/fiber/v2"
)

func ReadProduct(c *fiber.Ctx) error {
	var product []models.Product
	database.DB.First(&product)

	return c.Status(200).JSON(&fiber.Map{
		"status":  "Success",
		"payload": product,
	})
}
