package roles

import (
	"digital-cash-vault/internal/domain/common"
)

type Role struct {
	Id            common.RoleId
	Name          string
	PermissionIds []common.PermissionId
}
