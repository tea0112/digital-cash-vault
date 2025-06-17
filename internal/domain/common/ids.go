package common

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

type UserId string

func (id *UserId) IsZero() bool {
	return *id == ""
}

func (id *UserId) Scan(value any) error {
	switch v := value.(type) {
	case int64:
		*id = UserId(strconv.FormatInt(v, 10))
		return nil
	case []byte:
		*id = UserId(v)
		return nil
	case string:
		*id = UserId(v)
		return nil
	default:
		return fmt.Errorf("unsupported Scan type for UserId: %T", value)
	}
}

// Value converts UserId (string) -> DB value (int64 or string depending on schema)
func (id *UserId) Value() (driver.Value, error) {
	return string(*id), nil
}

type PermissionId string
type RoleId string
