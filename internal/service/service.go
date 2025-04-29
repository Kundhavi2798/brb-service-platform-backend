package service

//
//import (
//	"gorm.io/gorm"
//	_ "errors"
//)
//
//// Service struct now has a DB field
//type Service struct {
//	DB *gorm.DB
//}
//
//// UpdateServiceAvailability updates the service availability based on the given service ID and status
//func (s *Service) UpdateServiceAvailability(serviceID int, availability string) error {
//	// Start a transaction
//	tx := s.DB.Begin()
//	if tx.Error != nil {
//		return tx.Error
//	}
//
//	// Select the service from the database
//	var service struct {
//		ID       int
//		Status   string
//	}
//	err := tx.Model(&service).Where("id = ?", serviceID).First(&service).Error
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	// Update the service's availability
//	err = tx.Model(&service).Where("id = ?", serviceID).Update("availability", availability).Error
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	// Commit the transaction
//	return tx.Commit().Error
//}
