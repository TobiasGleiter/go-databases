package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrUserNotFound       = errors.New("models: no user found")
	ErrUserAlreadyExists  = errors.New("models: user already exists")
)
