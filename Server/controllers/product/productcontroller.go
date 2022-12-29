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

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	result := database.DB.Find(&product, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

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
		"status":  "Success",
		"payload": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	result := database.DB.Delete(&product, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&fiber.Map{
		"status": "Success",
	})
}
