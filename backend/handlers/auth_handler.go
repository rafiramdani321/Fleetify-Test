package handlers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	var userID int
	var role string

	if req.Username == "admin" && req.Password == "admin123" {
		userID = 1
		role = "Admin"
	} else if req.Username == "kerani" && req.Password == "kerani123" {
		userID = 2
		role = "Kerani"
	} else {
		return c.Status(401).JSON(fiber.Map{"message": "Wrong credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Could not login"})
	}

	return c.JSON(fiber.Map{"token": t, "role": role})
}