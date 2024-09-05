package repositories

import (
	"context"
	"socialmood/internal/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Create(ctx context.Context, userEntity *entities.UserEntity) (uuid.UUID, error) {
	query := `INSERT INTO users (email, password, first_name, last_name, birth, gender, interests, city) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var id uuid.UUID
	err := u.DB.QueryRow(ctx, query, userEntity.Email, userEntity.Password, userEntity.FirstName, userEntity.LastName, userEntity.Birth, userEntity.Gender, userEntity.Interests, userEntity.City).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (u *UserRepository) GetById(ctx context.Context, userId uuid.UUID) (entities.UserEntity, error) {
	query := `SELECT id, email, password, first_name, last_name, birth, gender, interests, city FROM users WHERE id = $1`
	var user entities.UserEntity
	err := u.DB.QueryRow(ctx, query, userId).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Birth, &user.Gender, &user.Interests, &user.City)
	if err != nil {
		return entities.UserEntity{}, err
	}

	return user, nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (entities.UserEntity, error) {
	query := `SELECT id, email, password, first_name, last_name, birth, gender, interests, city FROM users WHERE email = $1`
	var user entities.UserEntity
	err := u.DB.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Birth, &user.Gender, &user.Interests, &user.City)
	if err != nil {
		return entities.UserEntity{}, err
	}
	return user, nil
}
