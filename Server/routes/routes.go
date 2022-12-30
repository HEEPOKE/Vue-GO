package routes

import (
	productController "Server/Api/controllers/product"
	userController "Server/Api/controllers/user"
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

	api := r.Group("/api")

	user := api.Group("/user")
	user.Get("/get/:id", userController.GetUser)
	user.Get("/read", userController.ReadUser)
	user.Post("/add", userController.AddUser)
	user.Put("/update/:id", userController.UpdateUser)
	user.Delete("/delete/:id", userController.DeleteUser)

	product := api.Group("/product")
	product.Get("/get/:id", productController.GetProduct)
	product.Get("/read", productController.ReadProduct)
	product.Post("/add", productController.AddProduct)
	product.Put("/update/:id", productController.UpdateProduct)
	product.Delete("/delete/:id", productController.DeleteProduct)

	r.Listen(":6476")
}
