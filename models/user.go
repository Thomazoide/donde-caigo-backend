package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nombre             string
	Rut                string
	Email              string
	ProfilePicture     string
	ProfileDescription string
	Age                int
	Posts              []Post `gorm:"foreignKey:AuthorID"`
}
