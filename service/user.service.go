package service

import (
	"fmt"

	"github.com/Thomazoide/donde-caigo-backend/config"
	"github.com/Thomazoide/donde-caigo-backend/middleware"
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

func (s *UserService) CreateUser(nomrbe string, password string, rut string, email string, pfp string, desc string, age int64) (*models.User, error) {
	encrypter := middleware.NewEncrypter()
	if encrypter == nil {
		return nil, fmt.Errorf("failed to load enviroment variables")
	}
	hashPass, hashErr := encrypter.HashPassword(password)
	if hashErr != nil {
		return nil, hashErr
	}
	var newUser *models.User = models.CreateUser(nomrbe, hashPass, rut, email, pfp, desc, age)
	err := s.instance.Create(newUser)
	if err.Error != nil {
		return nil, err.Error
	}
	return newUser, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := s.instance.Select("ID", "nombre", "email", "rut", "profile_picture", "profile_description", "age").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user *models.User
	fmt.Println(id)
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
	result := s.instance.Where("ID = ?", usr.ID).Delete(usr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserService) UpdatePassword(id uint, actualPassword string, newPassword string) error {
	encrypter := middleware.NewEncrypter()
	var user *models.User
	err := s.instance.Where("ID = ?", id).Select("*").First(&user).Error
	if err != nil {
		return err
	}
	if !encrypter.VerifyPassword(actualPassword, user.Password) {
		return fmt.Errorf("clave incorrecta")
	}
	newHasehdPassword, hashErr := encrypter.HashPassword(newPassword)
	if hashErr != nil {
		return hashErr
	}
	user.Password = newHasehdPassword
	if saveErr := s.instance.Save(&user).Error; saveErr != nil {
		return saveErr
	}
	return nil
}
