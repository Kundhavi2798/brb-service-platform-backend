package service

import (
	"brb-service-platform-backend/pkg/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	testDB, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	testDB.AutoMigrate(&Category{}, &Service{}) // Only if needed
	db.DB = testDB
	return testDB
}

func TestCreateCategory(t *testing.T) {
	setupTestDB()

	category := &Category{Name: "Spa"}
	err := CreateCategory(category)

	assert.NoError(t, err)
	assert.NotZero(t, category.ID)
}

func TestCreateService(t *testing.T) {
	setupTestDB()

	service := &Service{Name: "Haircut", Price: 500, IsAvailable: true}
	err := CreateService(service)

	assert.NoError(t, err)
	assert.NotZero(t, service.ID)
}

func TestUpdateServiceAvailability(t *testing.T) {
	db := setupTestDB()
	service := &Service{Name: "Massage", Price: 1000, IsAvailable: true}
	db.Create(service)

	err := UpdateServiceAvailability(service.ID, false)

	assert.NoError(t, err)

	var updated Service
	db.First(&updated, service.ID)
	assert.False(t, updated.IsAvailable)
}

func TestUpdateServicePrice(t *testing.T) {
	db := setupTestDB()
	service := &Service{Name: "Facial", Price: 800, IsAvailable: true}
	db.Create(service)

	newPrice := 1200.0
	err := UpdateServicePrice(service.ID, newPrice)

	assert.NoError(t, err)

	var updated Service
	db.First(&updated, service.ID)
	assert.Equal(t, newPrice, updated.Price)
}
