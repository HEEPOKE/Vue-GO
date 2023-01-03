package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ValidationMiddleware(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "forbidden",
				"message": "Missing authorization header",
			})
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "forbidden",
				"message": "Invalid authorization header",
			})
		}
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("userId", claims)
		} else {
			c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "forbidden",
				"message": err.Error(),
			})
		}
		return c.Next()
	}
}
