package userUseCases

import (
	"context"
	"fmt"
	"socialmood/api/dto"
	"socialmood/internal/exceptions"
	"strings"

	"github.com/google/uuid"
)

func (u *UserUseCases) Login(ctx context.Context, loginDto dto.LoginDto) (string, error) {
	user, err := u.userRepository.GetByEmail(ctx, strings.ToLower(loginDto.Email))
	if err != nil {
		return "", err
	}
	if user.ID == uuid.Nil {
		return "", exceptions.ErrInvalidEmailOrPassword
	}

	err = user.ValidatePassword(loginDto.Password)
	if err != nil {
		fmt.Println(3)
		return "", exceptions.ErrInvalidEmailOrPassword
	}

	tokens, err := u.jwtInst.GenerateToken(user.GetID().String())
	if err != nil {
		return "", err
	}

	return tokens.AccessToken, nil
}
