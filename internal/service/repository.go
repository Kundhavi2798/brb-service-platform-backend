package service

import (
	"brb-service-platform-backend/pkg/db"
)

func CreateCategory(category *Category) error {
	return db.DB.Create(category).Error
}

func CreateService(service *Service) error {
	return db.DB.Create(service).Error
}

func UpdateServiceAvailability(serviceID uint, isAvailable bool) error {
	return db.DB.Model(&Service{}).Where("id = ?", serviceID).Update("is_available", isAvailable).Error
}

func UpdateServicePrice(serviceID uint, price float64) error {
	return db.DB.Model(&Service{}).Where("id = ?", serviceID).Update("price", price).Error
}
