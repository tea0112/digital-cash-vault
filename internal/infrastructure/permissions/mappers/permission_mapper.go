package mappers

import (
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/permissions"
	"digital-cash-vault/internal/infrastructure/permissions/models"
	"gorm.io/gorm"
	"strconv"
)

func MapEntityToDomain(e *models.PermissionEntity) *permissions.Permission {
	return &permissions.Permission{
		Id:   common.PermissionId(strconv.FormatUint(uint64(e.ID), 10)),
		Name: e.Name,
	}
}

func MapDomainToEntity(p *permissions.Permission) (*models.PermissionEntity, error) {
	id, err := strconv.Atoi(string(p.Id))
	if err != nil {
		return nil, err
	}

	return &models.PermissionEntity{
		Model: gorm.Model{
			ID: uint(id),
		},
		Name: p.Name,
	}, nil
}
