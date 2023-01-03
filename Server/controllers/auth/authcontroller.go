package auth

import (
	"Server/Api/database"
	"Server/Api/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	json := new(models.User)
	if err := c.BodyParser(json); err != nil {
		return err
	}
	if json.Username == "" || json.Email == "" || json.Password == "" || json.Tel == "" {
		return fiber.NewError(fiber.StatusBadRequest, "กรุณากรอกข้อมูลให้ครบ")
	}
	var userExist models.User
	database.DB.Where("email = ?", json.Email).First(&userExist)
	if userExist.ID > 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "Error",
			"message": "มีผู้ใช้งานอีเมล์นี้เเล้ว",
		})
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := models.User{
		Username: json.Username,
		Password: string(encryptedPassword),
		Email:    json.Email,
		Tel:      json.Tel,
		Role:     json.Role,
	}

	database.DB.Create(&user)
	if user.ID > 0 {
		c.Status(http.StatusOK).JSON(&fiber.Map{
			"status":  "ok",
			"message": "success",
			"userId":  user.ID,
		})
	} else {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "fail",
		})
	}
	return c.Status(http.StatusOK).JSON("")
}
