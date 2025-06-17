package repositories

import (
	"context"
	"digital-cash-vault/db"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/permissions"
)

type gormPermissionRepo struct{ db db.DB }

func NewPermissionRepository(db db.DB) permissions.PermissionRepository {
	return &gormPermissionRepo{db}
}

func (g gormPermissionRepo) FindByID(ctx context.Context, id common.PermissionId) (*permissions.Permission, error) {
	// fetch and map to domain

	//TODO implement me
	panic("implement me")
}

func (g gormPermissionRepo) Save(ctx context.Context, u *permissions.Permission) error {
	// map domain → entity and save

	//TODO implement me
	panic("implement me")
}
