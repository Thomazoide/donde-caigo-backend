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

type PostSchema struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Pics        string `json:"pics"`
	AuthorID    uint   `json:"author_id"`
	Stars       string `json:"stars"`
}

func (p *Post) ToSchema() *PostSchema {
	return &PostSchema{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Pics:        p.Pics,
		AuthorID:    p.AuthorID,
		Stars:       p.Stars,
	}
}

func CreatePost(title string, desc string, pics string, authorID uint) *Post {
	return &Post{
		Title:       title,
		Description: desc,
		Pics:        pics,
		AuthorID:    authorID,
	}
}
