package permissions

import (
	"digital-cash-vault/internal/domain/common"
)

type Permission struct {
	Id   common.PermissionId
	Name string
}
