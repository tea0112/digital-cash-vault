package users

import (
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

func NewUser(username string, email string, password string, firstName string, lastName string) User {
	return User{
		Username:  username,
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
}
