package notification

import (
	"brb-service-platform-backend/pkg/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Assuming your Notification model is like this:

func setupTestDB(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.DB = testDB

	err = db.DB.AutoMigrate(&Notification{})
	assert.NoError(t, err)
}

func TestCreateNotification(t *testing.T) {
	setupTestDB(t)

	notif := &Notification{
		Message: "New message",
		Status:  "pending",
	}

	err := CreateNotification(notif)
	assert.NoError(t, err)

	var fetched Notification
	err = db.DB.First(&fetched, notif.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "New message", fetched.Message)
}

func TestUpdateNotification(t *testing.T) {
	setupTestDB(t)

	notif := &Notification{Message: "Old", Status: "pending"}
	db.DB.Create(notif)

	notif.Message = "Updated message"
	notif.Status = "sent"
	err := UpdateNotification(notif)
	assert.NoError(t, err)

	var fetched Notification
	db.DB.First(&fetched, notif.ID)
	assert.Equal(t, "Updated message", fetched.Message)
	assert.Equal(t, "sent", fetched.Status)
}

func TestGetPendingNotifications(t *testing.T) {
	setupTestDB(t)

	// Seed test data
	db.DB.Create(&Notification{Message: "Msg 1", Status: "pending"})
	db.DB.Create(&Notification{Message: "Msg 2", Status: "pending"})
	db.DB.Create(&Notification{Message: "Msg 3", Status: "sent"})

	result, err := GetPendingNotifications(10)
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	for _, n := range result {
		assert.Equal(t, "pending", n.Status)
	}
}
