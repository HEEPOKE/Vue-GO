package user

import (
	"Server/Api/database"
	"Server/Api/models"

	"github.com/gofiber/fiber/v2"
)

func ReadUser(c *fiber.Ctx) error {
	var user []models.User
	database.DB.Find(&user)

	return c.Status(200).JSON(&fiber.Map{
		"status":  "Success",
		"payload": user,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	result := database.DB.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&fiber.Map{
		"status":  "Success",
		"payload": user,
	})
}

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Create(&user)

	return c.Status(200).JSON(&fiber.Map{
		"status":  "Success",
		"payload": user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(models.User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(&fiber.Map{
		"status":  "Success",
		"payload": user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	result := database.DB.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&fiber.Map{
		"status": "Success",
	})
}
