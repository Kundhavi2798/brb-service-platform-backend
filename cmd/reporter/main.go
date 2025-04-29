package main

import (
	"brb-service-platform-backend/config"
	"brb-service-platform-backend/internal/report"
	"brb-service-platform-backend/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	cfg := configs.LoadConfig()
	db.ConnectDatabase(cfg)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "DB down"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	report.RegisterReportRoutes(r)

	log.Println("📊 Report Service running on :8085")
	r.Run(":8085")
}
