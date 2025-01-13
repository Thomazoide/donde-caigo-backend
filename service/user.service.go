package service

import (
	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/models"
	"gorm.io/gorm"
)

type UserService struct {
	instance *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		instance: config.GetInstance(),
	}
}

func (s *UserService) CreateUser(nomrbe string, rut string, email string, pfp string, desc string, age int64) (*models.User, error) {
	var newUser *models.User = models.CreateUser(nomrbe, rut, email, pfp, desc, age)
	err := s.instance.Create(newUser)
	if err.Error != nil {
		return nil, err.Error
	}
	return newUser, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := s.instance.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user *models.User
	result := s.instance.Where("ID = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (s *UserService) UpdateUser(usr *models.User) (*models.User, error) {
	result := s.instance.Save(&usr)
	if result.Error != nil {
		return nil, result.Error
	}
	return usr, nil
}

func (s *UserService) DeleteUser(usr *models.User) error {
	result := s.instance.Delete(usr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
