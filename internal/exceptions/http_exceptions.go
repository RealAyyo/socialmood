package exceptions

import "errors"

var (
	ErrBadRequest = errors.New("bad Request")
	ErrForbidden  = errors.New("forbidden")
)
