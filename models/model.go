package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}
type Post struct {
	gorm.Model
	UserID  uint
	Title   string
	Content string
}
