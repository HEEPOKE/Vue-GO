package auth

import (
	"Server/Api/database"
	"Server/Api/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

var hmacSampleSecret []byte

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
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"status":  "ok",
			"message": "success",
			"userId":  user.ID,
		})
	} else {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "fail",
		})
	}
}

func Login(c *fiber.Ctx) error {
	json := new(models.User)
	if err := c.BodyParser(json); err != nil {
		return err
	}
	if json.Email == "" || json.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "กรุณากรอกข้อมูลให้ครบ")
	}
	var userExist models.User
	database.DB.Where("email = ?", json.Email).First(&userExist)
	if userExist.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "Error",
			"message": "user Does Not Exist",
		})
	}
	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExist.ID,
			"exp":    time.Now().Add(time.Minute * 1).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"status":  "ok",
			"message": "Login Success",
			"token":   tokenString,
		})

	} else {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "Login Failed",
		})

	}
}
