package notification

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	Message     string    `json:"message"`
	Status      string    `json:"status"` // pending, sent, failed
	RetryCount  int       `json:"retry_count"`
	LastAttempt time.Time `json:"last_attempt"`
}
