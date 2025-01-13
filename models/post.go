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

func CreatePost(title string, desc string, pics string, authorID uint, stars string) *Post {
	return &Post{
		Title:       title,
		Description: desc,
		Pics:        pics,
		AuthorID:    authorID,
		Stars:       stars,
	}
}
