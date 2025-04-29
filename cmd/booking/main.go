package main

import (
	"brb-service-platform-backend/config"
	"brb-service-platform-backend/internal/booking"
	"brb-service-platform-backend/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	cfg := configs.LoadConfig()
	db.ConnectDatabase(cfg)

	db.DB.AutoMigrate(&booking.Booking{})

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "DB down"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	booking.RegisterBookingRoutes(r)

	r.Run(":8082")
}
