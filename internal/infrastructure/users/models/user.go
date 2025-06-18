package models

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model

	Email    string `gorm:"uniqueIndex;size:255;not null"`
	Name     string `gorm:"size:120;not null"`
	Password string `gorm:"size:60;not null"`
}
