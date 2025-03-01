package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nombre             string
	Rut                string
	Email              string `gorm:"unique"`
	ProfilePicture     string
	ProfileDescription string
	Password           string
	Age                int64
	Posts              []Post `gorm:"foreignKey:AuthorID;default:null"`
}

type UserSchema struct {
	ID                 uint   `json:"id"`
	Nombre             string `json:"nombre"`
	Rut                string `json:"rut"`
	Email              string `json:"email"`
	ProfilePicture     string `json:"profilePicture"`
	ProfileDescription string `json:"profileDescription"`
	Age                int64  `json:"age"`
}

func (u *User) ToSchema() *UserSchema {
	return &UserSchema{
		ID:                 u.ID,
		Nombre:             u.Nombre,
		Rut:                u.Rut,
		Email:              u.Email,
		ProfilePicture:     u.ProfilePicture,
		ProfileDescription: u.ProfileDescription,
		Age:                u.Age,
	}
}

func CreateUser(nombre string, password string, rut string, email string, pfp string, desc string, age int64) *User {
	return &User{
		Nombre:             nombre,
		Rut:                rut,
		Email:              email,
		ProfilePicture:     pfp,
		ProfileDescription: desc,
		Password:           password,
		Age:                age,
	}
}
