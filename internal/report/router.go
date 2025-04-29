package report

import "github.com/gin-gonic/gin"

func RegisterReportRoutes(r *gin.Engine) {
	r.GET("/reports/vendor/:id", VendorReportHandler)
}
