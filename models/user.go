package models

import "gorm.io/gorm"

// User represents a user entity with fields for ID, name, email, and password
type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}