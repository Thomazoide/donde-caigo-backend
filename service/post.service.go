package service

import (
	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/models"
	"gorm.io/gorm"
)

type PostService struct {
	instance *gorm.DB
}

func NewPostService() *PostService {
	return &PostService{
		instance: config.GetInstance(),
	}
}

func (s *PostService) GetAllPost() ([]models.Post, error) {
	var posts []models.Post
	result := s.instance.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (s *PostService) CreatePost(title string, desc string, pics string, authorID uint, stars string) (*models.Post, error) {
	post := models.CreatePost(title, desc, pics, authorID, stars)
	result := s.instance.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}
