package booking

import (
	"brb-service-platform-backend/pkg/db"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func BookHandler(c *gin.Context) {
	var req Booking
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := BookSlot(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func RescheduleHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var body struct {
		NewSlot string `json:"new_slot"` // ISO time
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTime, err := time.Parse(time.RFC3339, body.NewSlot)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format"})
		return
	}

	if !isValidSlot(newTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid slot"})
		return
	}

	// Prevent overlapping
	var booking Booking
	if err := db.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	if IsSlotBooked(booking.VendorID, newTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Slot already booked"})
		return
	}

	err = UpdateBookingSlot(uint(id), newTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not reschedule"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking rescheduled"})
}

func CancelHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := UpdateBookingStatus(uint(id), "cancelled")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled"})
}

func UserBookingsHandler(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, _ := strconv.Atoi(userIDParam)

	bookings, err := GetUserBookings(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
