package routes

import (
	authController "Server/Api/controllers/auth"
	productController "Server/Api/controllers/product"
	userController "Server/Api/controllers/user"
	"Server/Api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router() {
	r := fiber.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           12,
	}))

	validationMiddleware := func(c *fiber.Ctx) error {
		return middleware.ValidationMiddleware(func(c *fiber.Ctx) error {
			return nil
		})(c)
	}

	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	user := api.Group("/user", validationMiddleware)
	user.Get("/get/:id", userController.GetUser)
	user.Get("/read", userController.ReadUser)
	user.Post("/add", userController.AddUser)
	user.Put("/update/:id", userController.UpdateUser)
	user.Delete("/delete/:id", userController.DeleteUser)

	product := api.Group("/product", validationMiddleware)
	product.Get("/get/:id", productController.GetProduct)
	product.Get("/read", productController.ReadProduct)
	product.Post("/add", productController.AddProduct)
	product.Put("/update/:id", productController.UpdateProduct)
	product.Delete("/delete/:id", productController.DeleteProduct)

	r.Listen(":6476")
}
