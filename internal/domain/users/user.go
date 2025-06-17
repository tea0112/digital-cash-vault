package users

import (
	"context"
	"digital-cash-vault/internal/domain/common"
)

type User struct {
	Id        common.UserId
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type UserRepository interface {
	FindByID(ctx context.Context, id common.UserId) (*User, error)
	Save(ctx context.Context, u *User) error
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}

func NewUser(username string, email string, password string, firstName string, lastName string) User {
	return User{
		Username:  username,
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
}
