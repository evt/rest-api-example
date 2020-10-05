package model

import (
	"errors"
)

// Common errors
var (
	ErrNotFound   = errors.New("resource not found")
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal error")
)
