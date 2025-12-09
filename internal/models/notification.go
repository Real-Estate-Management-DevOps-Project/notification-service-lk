package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationType string
type NotificationStatus string

const (
	TypeEmail NotificationType = "EMAIL"
	TypeSMS   NotificationType = "SMS"
	TypePush  NotificationType = "PUSH"
)

const (
	StatusPending NotificationStatus = "PENDING"
	StatusSent    NotificationStatus = "SENT"
	StatusFailed  NotificationStatus = "FAILED"
)

type Notification struct {
	ID        uuid.UUID          `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Recipient string             `gorm:"not null" json:"recipient" validate:"required"`
	Type      NotificationType   `gorm:"type:varchar(20);not null" json:"type" validate:"required,oneof=EMAIL SMS PUSH"`
	Subject   string             `gorm:"size:255" json:"subject"`
	Content   string             `gorm:"type:text;not null" json:"content" validate:"required"`
	Status    NotificationStatus `gorm:"type:varchar(20);default:'PENDING'" json:"status"`
	Error     string             `gorm:"type:text" json:"error,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Notification) TableName() string {
	return "notifications"
}
