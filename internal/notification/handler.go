package notification

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendNotificationHandler(c *gin.Context) {
	var n Notification
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n.Status = "pending"
	n.RetryCount = 0
	if err := CreateNotification(&n); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save notification"})
		return
	}

	go ProcessNotification(&n)

	c.JSON(http.StatusAccepted, gin.H{"message": "Notification enqueued"})
}
