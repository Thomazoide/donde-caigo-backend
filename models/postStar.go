package models

type PostStar struct {
	PostID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`
}
