package routes

import (
	productController "Server/Api/controllers/product"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router() {
	r := fiber.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ENDPOINT_URL"),
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           12,
	}))

	r.Get("/product", productController.ReadProduct)
	r.Post("/product/add", productController.AddProduct)

	r.Listen(":6476")
}
