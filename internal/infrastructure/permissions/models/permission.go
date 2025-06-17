package models

import "gorm.io/gorm"

type PermissionEntity struct {
	gorm.Model

	Name string `gorm:"uniqueIndex;not null"`
}

func (PermissionEntity) TableName() string {
	return "permissions"
}
