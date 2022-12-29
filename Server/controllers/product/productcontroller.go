package product

import (
	"Server/Api/database"
	"Server/Api/models"

	"github.com/gofiber/fiber/v2"
)

func ReadProduct(c *fiber.Ctx) error {
	var product []models.Product
	database.DB.Find(&product)

	return c.Status(200).JSON(&fiber.Map{
		"status":  "Success",
		"payload": product,
	})
}

func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Create(&product)

	return c.Status(200).JSON(&fiber.Map{
		"status":  "Success",
		"payload": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	id := c.Params("id")

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Where("id = ?", id).Updates(&product)
	return c.Status(200).JSON(&fiber.Map{
		"status":  "Update Success",
		"payload": product,
	})
}
