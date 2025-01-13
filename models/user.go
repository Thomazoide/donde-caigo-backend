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
