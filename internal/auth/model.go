package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"` // hashed
	Role     string `json:"role"`     // admin/user
}
