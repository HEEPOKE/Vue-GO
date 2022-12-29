package routes

import (
	productController "Server/Api/controllers/product"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	r := fiber.New()

	r.Get("/product", productController.ReadProduct)
	r.Post("/product/add", productController.AddProduct)

	r.Listen(":6476")
}
