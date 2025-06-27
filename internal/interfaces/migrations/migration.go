package migrations

import (
	"digital-cash-vault/db"
	permissionModels "digital-cash-vault/internal/infrastructure/permissions/models"
	roleModels "digital-cash-vault/internal/infrastructure/roles/models"
	userModels "digital-cash-vault/internal/infrastructure/users/models"
	"gorm.io/gorm"
	"log"
)

func AutoMigrate(db db.DB) {
	switch v := db.Raw().(type) {
	case *gorm.DB:
		err := v.AutoMigrate(
			&userModels.UserEntity{},
			&permissionModels.PermissionEntity{},
			&roleModels.RoleEntity{},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
