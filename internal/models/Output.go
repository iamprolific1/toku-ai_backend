package models

import "time"

type Output struct {
	ID uint `gorm:"primaryKey; autoIncrement"`
	UploadID uint `gorm:"index;not null"`
	Type string `gorm:"not null"`
	Content string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}