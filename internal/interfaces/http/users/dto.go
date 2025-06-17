package users

import (
	"digital-cash-vault/internal/applications/users"
	domainUser "digital-cash-vault/internal/domain/users"
)

type CreateUserRequestDTO struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func CreateUserRequestDTOToDomain(dto CreateUserRequestDTO) domainUser.User {
	return domainUser.NewUser(dto.Username, dto.Email, dto.Password, dto.FirstName, dto.LastName)
}

type UserResponseDTO struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FromUserDomainToUserResponseDTO(u domainUser.User) UserResponseDTO {
	return UserResponseDTO{
		ID:        string(u.Id),
		UserName:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

type LoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequestDTO) ToInput() users.LoginInput {
	return users.LoginInput{
		Email:    r.Email,
		Password: r.Password,
	}
}

type UserDTO struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LoginResponseDTO struct {
	User  UserDTO `json:"user"`
	Token string  `json:"token"`
}
