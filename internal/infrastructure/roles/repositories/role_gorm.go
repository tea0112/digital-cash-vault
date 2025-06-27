package repositories

import (
	"context"
	"digital-cash-vault/db"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/roles"
)

type gormRoleRepo struct{ db db.DB }

func NewRoleRepo(db db.DB) roles.RoleRepo { return &gormRoleRepo{db} }

func (g gormRoleRepo) FindByID(ctx context.Context, id common.RoleId) (*roles.Role, error) {
	// fetch and map to domain

	//TODO implement me
	panic("implement me")
}

func (g gormRoleRepo) Save(ctx context.Context, u *roles.Role) error {
	// map domain → entity and save

	//TODO implement me
	panic("implement me")
}
