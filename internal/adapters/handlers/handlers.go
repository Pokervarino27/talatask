package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pokervarino27/talatask/internal/ports"
)

type Handler struct {
	TaskAssignmentService ports.TaskAssignmentService
}

func NewHandler(task ports.TaskAssignmentService) *Handler {
	return &Handler{TaskAssignmentService: task}
}

func (h *Handler) AssignTasks(c *fiber.Ctx) error {
	assignments, err := h.TaskAssignmentService.AssignTask()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": true,
			"error":   err.Error(),
		})
	}
	println("Asignacioens: ", assignments)
	return c.JSON(assignments)
}

func (h *Handler) GenerateReport(c *fiber.Ctx) error {
	dateStr := c.Query("date")
	if dateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": true,
			"error":   "date parameter is required",
		})
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": true,
			"error":   "Invalid date format. Use YYYY-MM-DD",
		})
	}

	report, err := h.TaskAssignmentService.GenerateReport(date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": true,
			"error":   err.Error(),
		})
	}

	return c.JSON(report)
}

func (h *Handler) SetupRoutes(app *fiber.App) {
	app.Post("/assign-tasks", h.AssignTasks)
	app.Get("/report", h.GenerateReport)
}
