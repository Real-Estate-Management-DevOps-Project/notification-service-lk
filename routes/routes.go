package routes

import (
	"notification-service/internal/database"
	"notification-service/internal/handler"
	"notification-service/internal/repository"
	"notification-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Initialize dependencies
	repo := repository.NewNotificationRepository(database.GetDB())
	svc := service.NewNotificationService(repo)
	h := handler.NewNotificationHandler(svc)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "notification-service",
			"version": "1.0.0",
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	notifications := v1.Group("/notifications")
	notifications.Post("/send", h.Send)
	notifications.Get("/history", h.GetHistory)
	notifications.Get("/:id", h.GetByID)
}
