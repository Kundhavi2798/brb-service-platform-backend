package booking

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	VendorID  uint      `json:"vendor_id"`
	ServiceID uint      `json:"service_id"`
	SlotTime  time.Time `json:"slot_time"`
	Status    string    `json:"status"` // "pending", "confirmed", etc.
}
