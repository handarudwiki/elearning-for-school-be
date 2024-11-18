package commons

import "errors"

var (
	ErrCredentials    = errors.New("invalid credentials")
	ErrNotFound       = errors.New("data not found")
	ErrInternalServer = errors.New("internal server error")
	ErrConflict       = errors.New("data already exists")
)
