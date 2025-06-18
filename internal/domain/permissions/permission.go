package permissions

import (
	"context"
	"digital-cash-vault/internal/domain/common"
)

type Permission struct {
	Id   common.PermissionId
	Name string
}

type PermissionRepository interface {
	FindByID(ctx context.Context, id common.PermissionId) (*Permission, error)
	Save(ctx context.Context, u *Permission) error
}
