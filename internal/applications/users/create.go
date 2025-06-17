package users

import (
	"context"
	"digital-cash-vault/internal/domain/users"
)

func (s *UserServiceImpl) CreateUser(ctx context.Context, input users.User) error {
	err := s.userRepository.Save(ctx, &input)
	if err != nil {
		return err
	}

	return nil
}
