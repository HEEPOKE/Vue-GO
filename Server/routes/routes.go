package routes

import (
	productController "Server/Api/controllers/product"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	r := fiber.New()
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	r.Get("/product", productController.ReadProduct)
	r.Listen(":6476")
}
