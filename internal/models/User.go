package models

import (
	"time"
	// "gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey; autoIncrement"`
	Name string `gorm:"not null"`
	Email string `gorm:"uniqueIndex; not null"`
	Password string `gorm:"not null"`
	CreatedAt time.Time
	Uploads      []Upload `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
