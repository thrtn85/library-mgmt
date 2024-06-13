package models

import "gorm.io/gorm"

// Book represents a book entity with fields for ID, title, and author
type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
}