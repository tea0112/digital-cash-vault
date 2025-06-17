package users

import (
	"context"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
)

type UserServiceImpl struct {
	userRepository users.UserRepository
}

func NewUserService(repo users.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: repo,
	}
}

type UserService interface {
	CreateUser(ctx context.Context, input users.User) error
	GetUserById(ctx context.Context, id common.UserId) (*users.User, error)
	Login(ctx context.Context, loginInput LoginInput) (*LoginOutput, error)
}
