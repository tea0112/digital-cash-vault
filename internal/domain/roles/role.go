package roles

import (
	"context"
	"digital-cash-vault/internal/domain/common"
)

type Role struct {
	Id   common.RoleId
	Name string
}

type RoleRepo interface {
	FindByID(ctx context.Context, id common.RoleId) (*Role, error)
	Save(ctx context.Context, u *Role) error
}
