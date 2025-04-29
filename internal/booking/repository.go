package booking

import (
	"brb-service-platform-backend/pkg/db"
	"time"
)

func CreateBooking(b *Booking) error {
	return db.DB.Create(b).Error
}

func GetUserBookings(userID uint) ([]Booking, error) {
	var bookings []Booking
	err := db.DB.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func IsSlotBooked(vendorID uint, slot time.Time) bool {
	var count int64
	db.DB.Model(&Booking{}).
		Where("vendor_id = ? AND slot_time = ? AND status IN ?", vendorID, slot, []string{"pending", "confirmed", "in-progress"}).
		Count(&count)
	return count > 0
}

func UpdateBookingStatus(id uint, status string) error {
	return db.DB.Model(&Booking{}).Where("id = ?", id).Update("status", status).Error
}

func UpdateBookingSlot(id uint, newSlot time.Time) error {
	return db.DB.Model(&Booking{}).Where("id = ?", id).Update("slot_time", newSlot).Error
}
