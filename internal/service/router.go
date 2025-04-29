package service

import "github.com/gin-gonic/gin"

func RegisterServiceRoutes(r *gin.Engine) {
	serviceRoutes := r.Group("/services")
	{
		serviceRoutes.POST("/categories", CreateCategoryHandler)
		serviceRoutes.POST("/", CreateServiceHandler)
		serviceRoutes.PUT("/:id/availability", ToggleAvailabilityHandler)
		serviceRoutes.PUT("/:id/price", UpdatePriceHandler)
	}
}
