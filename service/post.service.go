package service

import (
	"strconv"
	"time"

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

func (s *PostService) CreatePost(title string, desc string, pics string, authorID uint) (*models.Post, error) {
	post := models.CreatePost(title, desc, pics, authorID)
	var userExists *models.User
	exists := s.instance.Where("ID = ?", authorID).Select("ID").First(&userExists).Error
	if exists != nil {
		return nil, exists
	}
	result := s.instance.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (s *PostService) EditPost(post *models.Post) (*models.Post, error) {
	result := s.instance.Save(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (s *PostService) AddLike(post *models.Post, id uint) (*models.Post, error) {
	sid := strconv.FormatUint(uint64(id), 10)
	tmpPost := post
	tmpPost.Stars = tmpPost.Stars + "," + sid
	if err := s.instance.Save(&tmpPost).Error; err != nil {
		return nil, err
	}
	return tmpPost, nil
}

func (s *PostService) DeletePost(postID uint) error {
	var post *models.Post
	result := s.instance.Where("id = ?", postID).First(&post)
	if result.Error != nil {
		return result.Error
	}
	result = s.instance.Delete(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *PostService) DeleteOldPosts() error {
	cutOffTime := time.Now().Add(-24 * time.Hour)
	var oldPosts []models.Post
	if err := s.instance.Where("created_at < ?", cutOffTime).Find(&oldPosts).Error; err != nil {
		return err
	}
	// for _, post := range oldPosts {
	// 	picsURLs := strings.Split(post.Pics, ",")
	// }
	return nil
}
