package userUseCases

import (
	"context"
	"socialmood/api/dto"
	"socialmood/internal/entities"
	"socialmood/internal/exceptions"

	"github.com/jackc/pgx/v5/pgconn"
)

func (u *UserUseCases) WebRegisterFlow(ctx context.Context, registerDto dto.RegisterDto) error {
	userEntity := entities.GetUserEntity()
	mappedUser, err := userEntity.ConvertRegisterDtoToModel(registerDto)
	if err != nil {
		return err
	}

	_, err = u.userRepository.Create(ctx, &mappedUser)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok && pqErr.Code == "23505" {
			return exceptions.ErrEmailAlreadyRegister
		}
	}

	return nil
}
