package roles

import (
	"context"
	"digital-cash-vault/internal/domain/common"
	domainRoles "digital-cash-vault/internal/domain/roles"
)

type RoleRepo interface {
	FindByID(ctx context.Context, id common.RoleId) (*domainRoles.Role, error)
	Save(ctx context.Context, u *domainRoles.Role) error
}
