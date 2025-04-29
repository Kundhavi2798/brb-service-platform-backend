package service

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `json:"name" gorm:"unique;not null"`
	Services []Service `gorm:"foreignKey:CategoryID"`
}

type Service struct {
	gorm.Model
	Name        string  `json:"name"`
	VendorID    uint    `json:"vendor_id"`
	CategoryID  uint    `json:"category_id"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"is_available"`
}
