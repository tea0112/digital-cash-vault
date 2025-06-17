package repositories

import (
	"context"
	"digital-cash-vault/internal/db"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
)

type gormUserRepo struct{ db db.DB }

func NewUserRepo(db db.DB) users.UserRepository { return &gormUserRepo{db} }

func (g gormUserRepo) FindByID(ctx context.Context, id common.UserId) (*users.User, error) {
	// fetch and map to domain

	//TODO implement me
	panic("implement me")
}

func (g gormUserRepo) Save(ctx context.Context, u *users.User) error {
	// map domain → entity and save

	//TODO implement me
	panic("implement me")
}
