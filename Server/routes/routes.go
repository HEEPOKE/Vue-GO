package routes

import "github.com/gofiber/fiber/v2"

func Router() {
	r := fiber.New()
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	r.Listen(":6476")
}
