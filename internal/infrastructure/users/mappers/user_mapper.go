package mappers

import (
	"digital-cash-vault/internal/domain/users"
	userModels "digital-cash-vault/internal/infrastructure/users/models"
)

//func MapEntityToDomain(e *models.PermissionEntity) *permissions.Permission {
//	return &permissions.Permission{
//		Id:   common.PermissionId(strconv.FormatUint(uint64(e.ID), 10)),
//		Name: e.Name,
//	}
//}

func MapDomainToEntity(u *users.User) (*userModels.UserEntity, error) {
	return &userModels.UserEntity{
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}, nil
}
