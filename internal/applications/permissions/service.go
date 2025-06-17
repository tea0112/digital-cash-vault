package permissions

import (
	"digital-cash-vault/internal/domain/permissions"
)

type PermissionServiceImpl struct {
	repo permissions.PermissionRepository
}

func NewPermissionService(repo permissions.PermissionRepository) *PermissionServiceImpl {
	return &PermissionServiceImpl{
		repo: repo,
	}
}
