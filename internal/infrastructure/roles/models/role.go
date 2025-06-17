package models

import "gorm.io/gorm"

type RoleEntity struct {
	gorm.Model

	Name string `gorm:"uniqueIndex;not null"`
}

func (RoleEntity) TableName() string {
	return "roles"
}
