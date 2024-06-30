package usecase

import "errors"

var (
	ErrUserAlreadyRegistered  = errors.New("user already registered")
	ErrUserInvalidCredentials = errors.New("invalid credentials")
)
