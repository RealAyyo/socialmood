package controllers

import (
	"context"
	"errors"
	"socialmood/api/dto"
	"socialmood/internal/entities"
	"socialmood/internal/exceptions"
	"socialmood/internal/validators"

	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

type UserUseCases interface {
	WebRegisterFlow(ctx context.Context, registerDto dto.RegisterDto) error
	Login(ctx context.Context, loginDto dto.LoginDto) (string, error)
	GetById(ctx context.Context, userID uuid.UUID, inputID string) (entities.UserEntity, error)
	Search(ctx context.Context, searchDto dto.SearchDto) ([]entities.UserEntity, error)
}

type UserController struct {
	userUseCases UserUseCases
}

func NewUserController(userUseCases UserUseCases) *UserController {
	return &UserController{
		userUseCases: userUseCases,
	}
}

func (a *UserController) Register(ctx *atreugo.RequestCtx) error {
	var registerDto dto.RegisterDto
	err := validators.ValidatePostQuery(ctx, &registerDto)
	if err != nil {
		return ctx.JSONResponse(&Response{
			Message: exceptions.ErrBadRequest.Error(),
			Result:  0,
		}, fasthttp.StatusBadRequest)
	}

	err = a.userUseCases.WebRegisterFlow(ctx, registerDto)
	if err != nil {
		code := fasthttp.StatusInternalServerError
		if errors.Is(err, exceptions.ErrEmailAlreadyRegister) {
			code = fasthttp.StatusConflict
		}

		return ctx.JSONResponse(&Response{
			Message: err.Error(),
			Result:  0,
		}, code)
	}
	return ctx.JSONResponse(&Response{
		Message: "Success",
		Result:  0,
	}, fasthttp.StatusCreated)
}
func (a *UserController) GetById(ctx *atreugo.RequestCtx) error {
	userId := ctx.UserValue("user").(uuid.UUID)
	id := ctx.UserValue("id").(string)

	user, err := a.userUseCases.GetById(ctx, userId, id)
	if err != nil {
		code := fasthttp.StatusInternalServerError
		if errors.Is(err, exceptions.ErrForbidden) {
			code = fasthttp.StatusForbidden
		}
		if errors.Is(err, exceptions.ErrBadRequest) {
			code = fasthttp.StatusBadRequest
		}
		return ctx.JSONResponse(&Response{
			Message: err.Error(),
			Result:  0,
		}, code)
	}
	return ctx.JSONResponse(&Response{
		Data: user,
	}, fasthttp.StatusOK)
}

func (a *UserController) Search(ctx *atreugo.RequestCtx) error {
	var searchDto dto.SearchDto
	err := validators.ValidatePostQuery(ctx, &searchDto)
	if err != nil {
		return ctx.JSONResponse(&Response{
			Message: err.Error(),
			Result:  0,
		}, fasthttp.StatusBadRequest)
	}
	users, err := a.userUseCases.Search(ctx, searchDto)
	if err != nil {
		return ctx.JSONResponse(&Response{
			Message: err.Error(),
			Result:  0,
		}, fasthttp.StatusInternalServerError)
	}

	return ctx.JSONResponse(&Response{
		Data: users,
	}, fasthttp.StatusOK)
	return nil
}
