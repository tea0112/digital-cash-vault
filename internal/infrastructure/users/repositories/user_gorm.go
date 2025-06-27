package repositories

import (
	"context"
	"digital-cash-vault/db"
	applicationUser "digital-cash-vault/internal/applications/users"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
	userMappers "digital-cash-vault/internal/infrastructure/users/mappers"
	"digital-cash-vault/internal/infrastructure/users/models"
	"digital-cash-vault/pkg"
	"fmt"
	"gorm.io/gorm"
)

type gormUserRepository struct{ db db.DB }

func NewUserRepository(db db.DB) applicationUser.UserRepository { return &gormUserRepository{db} }

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

func (g gormUserRepository) FindUsers(ctx context.Context, params applicationUser.ListUsersParams) ([]users.User, error) {
	queryBuilder := db.NewGormQueryBuilder(g.db.Raw().(*gorm.DB))
	if params.Username != nil {
		queryBuilder = queryBuilder.Where("username LIKE ?", fmt.Sprintf("%%%s%%", *params.Username))
	}

	if params.Status != nil {
		queryBuilder = queryBuilder.Where("status = ?", *params.Status)
	}

	if params.EmailLike != nil {
		queryBuilder = queryBuilder.Where("email LIKE ?", fmt.Sprintf("%%%s%%", *params.EmailLike))
	}

	if params.CreatedFrom != nil {
		queryBuilder = queryBuilder.Where("created_at >= ?", *params.CreatedFrom)
	}
	if params.CreatedTo != nil {
		queryBuilder = queryBuilder.Where("created_at <= ?", *params.CreatedTo)
	}

	if params.OrderBy != "" {
		queryBuilder = queryBuilder.OrderBy(params.OrderBy)
	}

	if params.Size > 0 {
		offset := (params.Page - 1) * params.Size
		queryBuilder = queryBuilder.Limit(params.Size).Offset(offset)
	}

	var userEntities []models.UserEntity
	err := queryBuilder.BuildNative().(*gorm.DB).WithContext(ctx).Find(&userEntities, "").Error
	if err != nil {
		return nil, err
	}

	return userMappers.MapEntitiesToDomains(userEntities), nil
}
