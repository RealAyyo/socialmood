package validators

import (
	"encoding/json"
	"socialmood/internal/exceptions"

	"github.com/go-playground/validator/v10"
	"github.com/savsgio/atreugo/v11"
)

var (
	Validate = validator.New()
)

func ValidatePostQuery(ctx *atreugo.RequestCtx, data interface{}) error {
	err := json.Unmarshal(ctx.PostBody(), &data)
	if err != nil {
		return exceptions.ErrBadRequest
	}

	err = Validate.Struct(data)
	if err != nil {
		return exceptions.ErrBadRequest
	}

	return nil
}
