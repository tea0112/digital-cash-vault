package users

import (
	"context"
	"digital-cash-vault/internal/domain/common"
)

type User struct {
	Id       common.UserId
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	FindByID(ctx context.Context, id common.UserId) (*User, error)
	Save(ctx context.Context, u *User) error
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
