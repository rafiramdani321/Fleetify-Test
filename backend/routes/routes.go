package routes

import (
	"fleetify-backend/handlers"
	"fleetify-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api")

	api.Post("/login", handlers.Login)
	// api.Get("/items")

	invoice := api.Group("/invoices", middleware.AuthRequired)
	invoice.Post("/", handlers.CreateInvoice)
}