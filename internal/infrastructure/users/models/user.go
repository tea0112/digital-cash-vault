package models

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model

	Username  string `gorm:"uniqueIndex;size:255;not null"`
	Email     string `gorm:"uniqueIndex;size:255;not null"`
	Password  string `gorm:"size:60;not null"`
	FirstName string `gorm:"size:120;not null"`
	LastName  string `gorm:"size:120;not null"`
}

func (UserEntity) TableName() string {
	return "users"
}
