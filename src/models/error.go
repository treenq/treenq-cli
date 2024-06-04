package models

import (
	"errors"
)

var (
	ErrContextAlreadyExists = errors.New("context with such name already exists")
	ErrContextNotFound      = errors.New("context with such name not found")
)
