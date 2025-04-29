package auth

import (
	"brb-service-platform-backend/pkg/db"
)

func CreateUser(user *User) error {
	return db.DB.Create(user).Error
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
