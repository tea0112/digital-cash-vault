package users

import (
	"context"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
)

type UserServiceImpl struct {
	repo users.UserRepository
}

func NewUserService(repo users.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

type UserService interface {
	CreateUser(ctx context.Context, input CreateUserInput) (*users.User, error)
	GetUserById(ctx context.Context, id common.UserId) (*users.User, error)
}
