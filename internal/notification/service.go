package notification

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func SimulateSend(n *Notification) bool {
	log.Printf("[SIMULATE] Sending to user %d: %s", n.UserID, n.Message)
	return n.RetryCount%4 != 0
}

func ProcessNotification(n *Notification) {
	n.LastAttempt = time.Now()

	success := SimulateSend(n)
	if success {
		n.Status = "sent"
		log.Printf("[SUCCESS] Notification %d sent.", n.ID)
	} else {
		n.RetryCount++
		delay := time.Duration(math.Pow(2, float64(n.RetryCount))) * time.Second
		log.Printf("[RETRY] Notification %d failed. Retrying in %v", n.ID, delay)
		time.Sleep(delay)
		n.Status = "pending"
	}

	UpdateNotification(n)
}
func PendingNotificationsHandler(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	notifications, err := GetPendingNotifications(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, notifications)
}
