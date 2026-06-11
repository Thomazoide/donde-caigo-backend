package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string
	Description string
	Pics        string
	AuthorID    uint
	Lat         float64
	Lng         float64
	StarsCount  int
	StarUsers   []User `gorm:"many2many:post_stars;"`
}

type PostSchema struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Pics        string  `json:"pics"`
	AuthorID    uint    `json:"author_id"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	StarsCount  int     `json:"stars"`
}

func (p *Post) ToSchema() *PostSchema {
	return &PostSchema{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Pics:        p.Pics,
		AuthorID:    p.AuthorID,
		StarsCount:  p.StarsCount,
	}
}

func CreatePost(title string, desc string, pics string, authorID uint) *Post {
	return &Post{
		Title:       title,
		Description: desc,
		Pics:        pics,
		AuthorID:    authorID,
		StarsCount:  0,
	}
}
