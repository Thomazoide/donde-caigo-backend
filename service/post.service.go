package service

import (
	"context"
	"strconv"
	"time"

	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/middleware"
	"github.com/Thomazoide/donde-caigo-backend/models"
	"github.com/Thomazoide/donde-caigo-backend/structs"
	"gorm.io/gorm"
)

type PostService struct {
	instance          *gorm.DB
	pictureMiddleware *middleware.PicturesMiddleware
}

func NewPostService() *PostService {
	return &PostService{
		instance:          config.GetInstance(),
		pictureMiddleware: middleware.NewPicturesMiddleware(),
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

func (s *PostService) CreatePost(payload structs.NewPostPayload, ctx context.Context) (*models.Post, error) {
	var userExists *models.User
	exists := s.instance.Where("ID = ?", payload.PostData.AuthorID).Select("ID").First(&userExists).Error
	if exists != nil {
		return nil, exists
	}
	post := models.CreatePost(payload.PostData.Title, payload.PostData.Description, "", payload.PostData.AuthorID)
	result := s.instance.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	var pics string
	for index, pic := range payload.Pics {
		//file name go by postID_AuthorID_index
		pic = strconv.FormatUint(uint64(post.ID), 10) + "_" + strconv.FormatUint(uint64(post.AuthorID), 10) + "_" + strconv.Itoa(index)
		//base64 to image
		img, convErr := s.pictureMiddleware.Base64ToImage(pic)
		if convErr != nil {
			return nil, convErr
		}
		//upload image
		var picUrl string
		var uploadErr error
		if picUrl, uploadErr = s.pictureMiddleware.UploadImage(ctx, pic, img); uploadErr != nil {
			return nil, uploadErr
		}
		pics = pics + picUrl + ","
	}
	post.Pics = pics
	result = s.instance.Save(&post)
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

func (s *PostService) AddLike(post *models.PostSchema, userId uint) (*models.Post, error) {
	formatedUserId := strconv.FormatUint(uint64(userId), 10)
	var tmpPost *models.Post
	if err := s.instance.Where("id = ?", post.ID).First(&tmpPost).Error; err != nil {
		return nil, err
	}
	tmpPost.Stars = tmpPost.Stars + "," + formatedUserId
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
	return nil
}
