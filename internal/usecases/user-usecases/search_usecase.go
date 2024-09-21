package userUseCases

import (
	"context"
	"socialmood/api/dto"
	"socialmood/internal/entities"
)

func (u *UserUseCases) Search(ctx context.Context, searchDto dto.SearchDto) ([]entities.UserEntity, error) {
	users, err := u.userRepository.Search(ctx, searchDto)
	if err != nil {
		return nil, err
	}
	return users, nil
}
