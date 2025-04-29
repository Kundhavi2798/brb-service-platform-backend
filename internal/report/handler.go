package report

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func VendorReportHandler(c *gin.Context) {
	idParam := c.Param("id")
	vendorID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendor ID"})
		return
	}

	report, err := GetVendorReport(uint(vendorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate report"})
		return
	}

	c.JSON(http.StatusOK, report)
}
