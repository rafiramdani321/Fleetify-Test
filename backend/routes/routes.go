package routes

import (
	"fleetify-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api")

	api.Post("/login")
	api.Get("/items")

	invoice := api.Group("/invoices", middleware.AuthRequired)
	invoice.Post("/")
}