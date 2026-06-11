package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	PostID      *uint          `gorm:"index;default:NULL" json:"post_id,omitempty"`
	UserID      *uint          `gorm:"index;default:NULL" json:"user_id,omitempty"`
	ObjectName  string         `gorm:"size:512;not null" json:"object_name"` // path en el bucket
	Bucket      string         `gorm:"size:255;not null" json:"bucket"`
	Filename    string         `gorm:"size:255" json:"filename"`
	ContentType string         `gorm:"size:100" json:"content_type"`
	Size        int64          `json:"size"`
	Width       int            `json:"width,omitempty"`
	Height      int            `json:"height,omitempty"`
	ThumbName   string         `gorm:"size:512" json:"thumb_name,omitempty"`
	Checksum    string         `gorm:"size:128;index" json:"checksum,omitempty"`
	Status      string         `gorm:"size:20;index;default:'TMP'" json:"status"` // TMP | FINAL
	IsPrimary   bool           `gorm:"default:false" json:"is_primary"`
	Metadata    datatypes.JSON `json:"metadata,omitempty"`
}
