package users

import (
	"context"
	"digital-cash-vault/internal/domain/users"
)

type CreateUserInput struct {
	Email    string
	Name     string
	Password string
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, input CreateUserInput) (*users.User, error) {
	user := users.NewUser(input.Name, input.Email, input.Password)
	err := s.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
