package customerrors

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)
