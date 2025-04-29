package notification

import (
	"brb-service-platform-backend/pkg/db"
)

func CreateNotification(n *Notification) error {
	return db.DB.Create(n).Error
}

func UpdateNotification(n *Notification) error {
	return db.DB.Save(n).Error
}

func GetPendingNotifications(limit int) ([]Notification, error) {
	var notifications []Notification
	err := db.DB.Where("status = ?", "pending").Limit(limit).Find(&notifications).Error
	return notifications, err
}
