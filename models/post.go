package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string
	Description string
	Pics        string
	AuthorID    uint
	Stars       string `gorm:"default:null"`
}
