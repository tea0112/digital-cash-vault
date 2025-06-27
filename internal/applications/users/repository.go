package users

import (
	"context"
	"digital-cash-vault/internal/domain/common"
	users2 "digital-cash-vault/internal/domain/users"
)

type UserRepository interface {
	FindByID(ctx context.Context, id common.UserId) (*users2.User, error)
	Save(ctx context.Context, u *users2.User) error
	FindByUsername(ctx context.Context, username string) (*users2.User, error)
	FindByEmail(ctx context.Context, email string) (*users2.User, error)
	FindUsers(ctx context.Context, params ListUsersParams) ([]users2.User, error)
}
