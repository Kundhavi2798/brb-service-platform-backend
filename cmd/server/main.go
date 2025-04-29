package main

import (
	"brb-service-platform-backend/config"
	serviceRoutes "brb-service-platform-backend/internal/service"
	"brb-service-platform-backend/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	cfg := configs.LoadConfig()
	db.ConnectDatabase(cfg)

	db.DB.AutoMigrate(&serviceRoutes.Category{}, &serviceRoutes.Service{})

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "DB down"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	serviceRoutes.RegisterServiceRoutes(r)

	r.Run(":8081")
}
