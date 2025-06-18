package models

import "gorm.io/gorm"

type PermissionEntity struct {
	gorm.Model

	Name string `gorm:"uniqueIndex;not null"`
}
