package users

import (
	"context"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
	"digital-cash-vault/pkg"
	"digital-cash-vault/pkg/customerrors"
	"digital-cash-vault/pkg/jwt"
)

func (s *UserServiceImpl) GetUserById(ctx context.Context, id common.UserId) (*users.User, error) {
	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	User  *users.User
	Token string
}

func (s *UserServiceImpl) Login(ctx context.Context, loginInput LoginInput) (*LoginOutput, error) {
	user, err := s.userRepository.FindByEmail(ctx, loginInput.Email)
	if err != nil || !pkg.VerifyPassword(loginInput.Password, user.Password) {
		return nil, customerrors.ErrInvalidCredentials
	}

	token, err := jwt.GenerateJWT(user.Email)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		User:  user,
		Token: token,
	}, nil
}
