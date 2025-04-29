package booking

import "github.com/gin-gonic/gin"

func RegisterBookingRoutes(r *gin.Engine) {
	bookingGroup := r.Group("/bookings")
	{
		bookingGroup.POST("/", BookHandler)
		bookingGroup.PUT("/:id/reschedule", RescheduleHandler)
		bookingGroup.PUT("/:id/cancel", CancelHandler)
		bookingGroup.GET("/user/:id", UserBookingsHandler)
	}
}
