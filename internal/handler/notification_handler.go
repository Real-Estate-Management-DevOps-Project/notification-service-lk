package handler

import (
	"log"
	"notification-service/internal/models"
	"notification-service/internal/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type NotificationHandler struct {
	service   service.NotificationService
	validator *validator.Validate
}

func NewNotificationHandler(service service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *NotificationHandler) Send(c *fiber.Ctx) error {
	var notification models.Notification
	if err := c.BodyParser(&notification); err != nil {
		log.Printf("Failed to parse notification request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	log.Printf("Received notification: recipient=%s, type=%s, subject=%s, content=%s", notification.Recipient, notification.Type, notification.Subject, notification.Content)

	if err := h.validator.Struct(notification); err != nil {
		log.Printf("Notification validation failed: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("Received notification request: type=%s, recipient=%s, subject=%s", notification.Type, notification.Recipient, notification.Subject)

	if err := h.service.SendNotification(&notification); err != nil {
		log.Printf("Failed to queue notification: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send notification"})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Notification queued for sending",
		"id":      notification.ID,
	})
}

func (h *NotificationHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	notification, err := h.service.GetNotification(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Notification not found"})
	}

	return c.JSON(notification)
}

func (h *NotificationHandler) GetHistory(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	notifications, total, err := h.service.GetHistory(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch history"})
	}

	return c.JSON(fiber.Map{
		"data": notifications,
		"meta": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}
