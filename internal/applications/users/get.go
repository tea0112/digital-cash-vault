package users

import (
	"context"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
)

func (s *UserServiceImpl) GetUserById(ctx context.Context, id common.UserId) (*users.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
