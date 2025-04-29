package main

import (
	"brb-service-platform-backend/config"
	"brb-service-platform-backend/internal/notification"
	"brb-service-platform-backend/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	cfg := configs.LoadConfig()
	db.ConnectDatabase(cfg)

	db.DB.AutoMigrate(&notification.Notification{})

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "DB down"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	notification.RegisterNotificationRoutes(r)

	log.Println("🚀 Notification Service running on :8083")
	r.Run(":8083")
}
