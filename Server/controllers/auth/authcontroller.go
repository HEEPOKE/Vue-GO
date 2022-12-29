package auth

import (
	"Server/Api/database"
	"Server/Api/models"

	"github.com/gofiber/fiber/v2"
)

func ReadProduct(c *fiber.Ctx) {
	var product []models.Product
	database.DB.First(&product)
	c.JSON(fiber.StatusOK, product)
}
