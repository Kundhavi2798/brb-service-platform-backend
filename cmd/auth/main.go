package main

import (
	"brb-service-platform-backend/config"
	"brb-service-platform-backend/internal/auth"
	"brb-service-platform-backend/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	cfg := configs.LoadConfig()
	db.ConnectDatabase(cfg)

	db.DB.AutoMigrate(&auth.User{})

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "DB down"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	auth.RegisterAuthRoutes(r)

	log.Println("🔐 Auth Service running on :8084")
	r.Run(":8084")
}
