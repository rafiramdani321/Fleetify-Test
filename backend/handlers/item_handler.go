package handlers

import (
	"fleetify-backend/config"
	"fleetify-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetItems(c *fiber.Ctx) error {
	var items []models.Item

	search := c.Query("search")

	db := config.DB

	if search != "" {
		searchTerm := "%" + search + "%"
		db = db.Where("name ILIKE ? OR code ILIKE ?", searchTerm, searchTerm)
	}

	if err := db.Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to get item"})
	}

	return c.JSON(items)
}