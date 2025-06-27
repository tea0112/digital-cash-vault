package mappers

import (
	"digital-cash-vault/internal/domain/common"
	"digital-cash-vault/internal/domain/roles"
	"digital-cash-vault/internal/infrastructure/authmodels"
	"strconv"
)

func MapAuthRoleEntityToDomain(e *authmodels.Role) *roles.Role {
	permissionIds := make([]common.PermissionId, len(e.Permissions))
	for i, permission := range e.Permissions {
		permissionIds[i] = common.PermissionId(strconv.Itoa(int(permission.Id)))
	}

	return &roles.Role{
		Id:            common.RoleId(strconv.Itoa(int(e.Id))),
		Name:          e.Name,
		PermissionIds: permissionIds,
	}
}

func MapDomainToEntity() {
}
