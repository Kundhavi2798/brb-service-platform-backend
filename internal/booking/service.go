package booking

import (
	"errors"
	"time"
)

var allowedHours = []int{9, 10, 11, 12, 13, 14, 15, 16} // 9 AM to 5 PM exclusive

func isValidSlot(slot time.Time) bool {
	hour := slot.Hour()
	for _, h := range allowedHours {
		if hour == h && slot.Minute() == 0 {
			return true
		}
	}
	return false
}

func BookSlot(b *Booking) error {
	if !isValidSlot(b.SlotTime) {
		return errors.New("invalid time slot, only full hours between 9 AM and 5 PM allowed")
	}

	if IsSlotBooked(b.VendorID, b.SlotTime) {
		return errors.New("slot already booked")
	}

	b.Status = "pending"
	return CreateBooking(b)
}
