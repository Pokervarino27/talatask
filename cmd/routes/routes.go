package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pokervarino27/talatask/internal/adapters/handlers"
)

func SetupRoutes(app *fiber.App, handler *handlers.Handler) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server OK!")
	})

	api := app.Group("/api/v1")

	api.Get("/report", handler.GenerateReport)
	api.Post("/assign-tasks", handler.AssignTasks)
}
