package controllers

import (
	"errors"
	"socialmood/api/dto"
	"socialmood/internal/exceptions"
	"socialmood/internal/validators"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

type AuthController struct {
	userUseCases UserUseCases
}

func NewAuthController(userUseCases UserUseCases) *AuthController {
	return &AuthController{
		userUseCases: userUseCases,
	}
}

func (a *AuthController) Login(ctx *atreugo.RequestCtx) error {
	var loginDto dto.LoginDto
	err := validators.ValidatePostQuery(ctx, &loginDto)
	if err != nil {
		return ctx.JSONResponse(&Response{
			Message: exceptions.ErrBadRequest.Error(),
			Result:  0,
		}, fasthttp.StatusBadRequest)
	}

	accessToken, err := a.userUseCases.Login(ctx, loginDto)
	if err != nil {
		code := fasthttp.StatusInternalServerError
		if errors.Is(err, exceptions.ErrInvalidEmailOrPassword) {
			code = fasthttp.StatusNotFound
		}

		return ctx.JSONResponse(&Response{
			Message: err.Error(),
			Result:  0,
		}, code)
	}
	return ctx.JSONResponse(&Response{
		Data: struct {
			AccessToken string `json:"access_token"`
		}{
			AccessToken: accessToken,
		},
	}, fasthttp.StatusOK)
}
