package repository

import (
	"notification-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(notification *models.Notification) error
	FindByID(id uuid.UUID) (*models.Notification, error)
	FindAll(offset, limit int) ([]models.Notification, int64, error)
	UpdateStatus(id uuid.UUID, status models.NotificationStatus, errorMsg string) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) Create(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) FindByID(id uuid.UUID) (*models.Notification, error) {
	var notification models.Notification
	if err := r.db.First(&notification, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *notificationRepository) FindAll(offset, limit int) ([]models.Notification, int64, error) {
	var notifications []models.Notification
	var total int64

	r.db.Model(&models.Notification{}).Count(&total)

	err := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&notifications).Error
	return notifications, total, err
}

func (r *notificationRepository) UpdateStatus(id uuid.UUID, status models.NotificationStatus, errorMsg string) error {
	return r.db.Model(&models.Notification{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status": status,
		"error":  errorMsg,
	}).Error
}
