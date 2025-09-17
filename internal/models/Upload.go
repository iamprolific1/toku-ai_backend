package models

import "time"

type Upload struct {
	ID uint `gorm:"primaryKey; autoIncrement"`
	UserID uint `gorm:"index;not null"`
	FilePath string `gorm:"type:text;not null"`
	Status string `gorm:"default:'pending'"`
	Transcript string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Outputs      []Output `gorm:"foreignKey:UploadID;constraint:OnDelete:CASCADE;"`
}