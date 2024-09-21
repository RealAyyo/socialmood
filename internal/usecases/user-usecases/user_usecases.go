package userUseCases

import (
	"context"
	"socialmood/api/dto"
	"socialmood/internal/config"
	"socialmood/internal/entities"
	"socialmood/internal/exceptions"
	"socialmood/internal/jwt"

	"github.com/google/uuid"
)

type UserUseCases struct {
	jwtInst        *jwt.JWT
	userRepository UserRepository
}

type UserRepository interface {
	Create(ctx context.Context, userEntity *entities.UserEntity) (uuid.UUID, error)
	GetById(ctx context.Context, userId uuid.UUID) (entities.UserEntity, error)
	GetByEmail(ctx context.Context, email string) (entities.UserEntity, error)
	Search(ctx context.Context, searchDto dto.SearchDto) ([]entities.UserEntity, error)
}

func NewUserUseCases(userRepository UserRepository, jwtConf *config.JWTConf) *UserUseCases {
	return &UserUseCases{
		jwtInst:        jwt.New(jwtConf),
		userRepository: userRepository,
	}
}

func (u *UserUseCases) GetById(ctx context.Context, userID uuid.UUID, inputID string) (entities.UserEntity, error) {
	parsedInputID, err := uuid.Parse(inputID)
	if err != nil {
		return entities.UserEntity{}, exceptions.ErrBadRequest
	}

	if parsedInputID != userID {
		return entities.UserEntity{}, exceptions.ErrForbidden
	}

	user, err := u.userRepository.GetById(ctx, parsedInputID)
	user.Password = ""

	return user, err
}
