package mappers

import (
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/users"
	userModels "digital-cash-vault/internal/infrastructure/users/models"
	"strconv"
)

func MapEntityToDomain(e userModels.UserEntity) users.User {
	return users.User{
		Id:        common.UserId(strconv.Itoa(int(e.ID))),
		Username:  e.Username,
		Email:     e.Email,
		Password:  e.Password,
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}
}

func MapEntitiesToDomains(entities []userModels.UserEntity) []users.User {
	domainUsers := make([]users.User, len(entities))
	for i, e := range entities {
		domainUsers[i] = MapEntityToDomain(e)
	}

	return domainUsers
}

func MapDomainToEntity(u *users.User) (*userModels.UserEntity, error) {
	return &userModels.UserEntity{
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}, nil
}
