package common

type UserId string
type PermissionId string
type RoleId string

func (id UserId) IsZero() bool {
	return id == ""
}
