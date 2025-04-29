package report

import (
	"brb-service-platform-backend/pkg/db"
)

type Report struct {
	TotalBookings  int     `json:"total_bookings"`
	TotalRevenue   float64 `json:"total_revenue"`
	TopServiceName string  `json:"top_service_name"`
}

func GetVendorReport(vendorID uint) (*Report, error) {
	var report Report

	// Total bookings
	err := db.DB.Raw(`
		SELECT COUNT(*) 
		FROM bookings 
		WHERE vendor_id = ? AND status IN ('pending', 'confirmed', 'in-progress', 'completed')
	`, vendorID).Scan(&report.TotalBookings).Error
	if err != nil {
		return nil, err
	}

	// Total revenue
	err = db.DB.Raw(`
		SELECT COALESCE(SUM(s.price), 0) 
		FROM bookings b
		JOIN services s ON b.service_id = s.id
		WHERE b.vendor_id = ? AND b.status IN ('completed')
	`, vendorID).Scan(&report.TotalRevenue).Error
	if err != nil {
		return nil, err
	}

	// Most frequently booked service
	err = db.DB.Raw(`
		SELECT s.name 
		FROM bookings b
		JOIN services s ON b.service_id = s.id
		WHERE b.vendor_id = ? AND b.status IN ('completed')
		GROUP BY s.name 
		ORDER BY COUNT(*) DESC 
		LIMIT 1
	`, vendorID).Scan(&report.TopServiceName).Error
	if err != nil {
		return nil, err
	}

	return &report, nil
}
