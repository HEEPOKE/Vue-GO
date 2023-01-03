package auth

import (
	"Server/Api/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var json models.User
	if err := c.
}
