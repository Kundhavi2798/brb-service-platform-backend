package booking

import (
	"brb-service-platform-backend/pkg/db"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.DB = testDB

	err = db.DB.AutoMigrate(&Booking{})
	assert.NoError(t, err)
}

func TestCreateBooking(t *testing.T) {
	setupTestDB(t)

	booking := &Booking{
		UserID:   1,
		VendorID: 2,
		SlotTime: time.Now(),
		Status:   "pending",
	}

	err := CreateBooking(booking)
	assert.NoError(t, err)

	var found Booking
	err = db.DB.First(&found, booking.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, booking.Status, found.Status)
}

func TestGetUserBookings(t *testing.T) {
	setupTestDB(t)

	db.DB.Create(&Booking{UserID: 1, VendorID: 2, SlotTime: time.Now(), Status: "pending"})
	db.DB.Create(&Booking{UserID: 1, VendorID: 3, SlotTime: time.Now(), Status: "confirmed"})
	db.DB.Create(&Booking{UserID: 2, VendorID: 3, SlotTime: time.Now(), Status: "completed"})

	bookings, err := GetUserBookings(1)
	assert.NoError(t, err)
	assert.Len(t, bookings, 2)
	for _, b := range bookings {
		assert.Equal(t, uint(1), b.UserID)
	}
}

func TestIsSlotBooked(t *testing.T) {
	setupTestDB(t)

	slot := time.Now()
	db.DB.Create(&Booking{UserID: 1, VendorID: 10, SlotTime: slot, Status: "pending"})
	db.DB.Create(&Booking{UserID: 2, VendorID: 10, SlotTime: slot, Status: "completed"}) // should be ignored

	isBooked := IsSlotBooked(10, slot)
	assert.True(t, isBooked)

	isBooked = IsSlotBooked(99, slot)
	assert.False(t, isBooked)
}

func TestUpdateBookingStatus(t *testing.T) {
	setupTestDB(t)

	b := &Booking{UserID: 1, VendorID: 2, SlotTime: time.Now(), Status: "pending"}
	db.DB.Create(b)

	err := UpdateBookingStatus(b.ID, "completed")
	assert.NoError(t, err)

	var updated Booking
	db.DB.First(&updated, b.ID)
	assert.Equal(t, "completed", updated.Status)
}

func TestUpdateBookingSlot(t *testing.T) {
	setupTestDB(t)

	oldSlot := time.Now()
	newSlot := oldSlot.Add(1 * time.Hour)

	b := &Booking{UserID: 1, VendorID: 2, SlotTime: oldSlot, Status: "confirmed"}
	db.DB.Create(b)

	err := UpdateBookingSlot(b.ID, newSlot)
	assert.NoError(t, err)

	var updated Booking
	db.DB.First(&updated, b.ID)
	assert.WithinDuration(t, newSlot, updated.SlotTime, time.Second)
}
