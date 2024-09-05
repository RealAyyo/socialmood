package exceptions

import "errors"

var (
	ErrInvalidBirthDate       = errors.New("invalid birth date format")
	ErrInvalidGender          = errors.New("invalid gender")
	ErrEmailAlreadyRegister   = errors.New("email already register")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)
