package dto

type RegisterDto struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Birth     string `json:"birth" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	Interests string `json:"interests" validate:"required"`
	City      string `json:"city" validate:"required"`
}
