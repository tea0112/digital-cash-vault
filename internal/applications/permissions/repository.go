package permissions

import (
	"context"
	"digital-cash-vault/internal/domain/common"
	permissions2 "digital-cash-vault/internal/domain/permissions"
)

type PermissionRepository interface {
	FindByID(ctx context.Context, id common.PermissionId) (*permissions2.Permission, error)
	Save(ctx context.Context, u *permissions2.Permission) error
}
