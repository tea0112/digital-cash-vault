package repositories

import (
	"context"
	"digital-cash-vault/db"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
	userMappers "digital-cash-vault/internal/infrastructure/users/mappers"
	"digital-cash-vault/pkg"
	"gorm.io/gorm"
)

type gormUserRepository struct{ db db.DB }

func NewUserRepository(db db.DB) users.UserRepository { return &gormUserRepository{db} }

func (g gormUserRepository) FindByID(ctx context.Context, id common.UserId) (*users.User, error) {
	// fetch and map to domain

	//TODO implement me
	panic("implement me")
}

func (g gormUserRepository) Save(ctx context.Context, u *users.User) error {
	user, err := userMappers.MapDomainToEntity(u)
	if err != nil {
		return err
	}

	hashedPassword, err := pkg.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return g.db.Tx(ctx, func(tx db.DB) error {
		gormDbTx := tx.Raw().(*gorm.DB)

		gormDbTx = gormDbTx.Create(user)
		if gormDbTx.Error != nil {
			return gormDbTx.Error
		}

		return nil
	})
}

func (g gormUserRepository) FindByUsername(ctx context.Context, username string) (*users.User, error) {
	panic("implement me")
}

func (g gormUserRepository) FindByEmail(ctx context.Context, email string) (*users.User, error) {
	var user users.User
	err := g.db.First(ctx, &user, "email = ?", email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
