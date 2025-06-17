package users

import (
	"context"
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
	"digital-cash-vault/pkg"
	"digital-cash-vault/pkg/customerrors"
	"digital-cash-vault/pkg/jwt"
	"time"
)

type UserServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(repo UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: repo,
	}
}

type UserService interface {
	CreateUser(ctx context.Context, input users.User) error
	GetUserById(ctx context.Context, id common.UserId) (*users.User, error)
	Login(ctx context.Context, loginInput LoginInput) (*LoginOutput, error)
	ListUsers(ctx context.Context, params ListUsersParams) ([]users.User, error)
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, input users.User) error {
	err := s.userRepository.Save(ctx, &input)
	if err != nil {
		return err
	}

	return nil
}

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

type ListUsersParams struct {
	Username    *string
	Status      *string
	EmailLike   *string
	CreatedFrom *time.Time
	CreatedTo   *time.Time
	OrderBy     string
	Page        int
	Size        int
}

func (s *UserServiceImpl) ListUsers(ctx context.Context, params ListUsersParams) ([]users.User, error) {
	return s.userRepository.FindUsers(ctx, params)
}
