package notification

import "github.com/gin-gonic/gin"

func RegisterNotificationRoutes(r *gin.Engine) {
	r.POST("/notify", SendNotificationHandler)
}
