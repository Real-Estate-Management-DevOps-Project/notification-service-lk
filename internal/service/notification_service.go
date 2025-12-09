package service

import (
	"log"
	"notification-service/internal/models"
	"notification-service/internal/repository"

	"github.com/google/uuid"
)

type NotificationService interface {
	SendNotification(notification *models.Notification) error
	GetNotification(id uuid.UUID) (*models.Notification, error)
	GetHistory(page, limit int) ([]models.Notification, int64, error)
}

type notificationService struct {
	repo repository.NotificationRepository
}

func NewNotificationService(repo repository.NotificationRepository) NotificationService {
	return &notificationService{repo: repo}
}

func (s *notificationService) SendNotification(notification *models.Notification) error {
	// First save as PENDING
	notification.Status = models.StatusPending
	if err := s.repo.Create(notification); err != nil {
		return err
	}

	// Mock sending process
	// In a real app, we would use an email/SMS provider here.
	// For now, we simulate success.

	go func(n *models.Notification) {
		// Simulate network delay
		// time.Sleep(100 * time.Millisecond)

		log.Printf("Sending %s to %s: %s", n.Type, n.Recipient, n.Subject)

		// Update status to SENT
		if err := s.repo.UpdateStatus(n.ID, models.StatusSent, ""); err != nil {
			log.Printf("Failed to update notification status: %v", err)
		}
	}(notification)

	return nil
}

func (s *notificationService) GetNotification(id uuid.UUID) (*models.Notification, error) {
	return s.repo.FindByID(id)
}

func (s *notificationService) GetHistory(page, limit int) ([]models.Notification, int64, error) {
	offset := (page - 1) * limit
	return s.repo.FindAll(offset, limit)
}
