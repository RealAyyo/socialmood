package entities

import (
	"socialmood/api/dto"
	"socialmood/internal/exceptions"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birth     time.Time `json:"birth"`
	Gender    string    `json:"gender"`
	Interests string    `json:"interests"`
	City      string    `json:"city"`
}

func GetUserEntity() *UserEntity {
	return &UserEntity{}
}

func (u *UserEntity) ConvertRegisterDtoToModel(registerDto dto.RegisterDto) (UserEntity, error) {
	birthDate, err := time.Parse("2006-01-02", registerDto.Birth)
	if err != nil {
		return UserEntity{}, exceptions.ErrInvalidBirthDate
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserEntity{}, err
	}

	return UserEntity{
		FirstName: registerDto.FirstName,
		LastName:  registerDto.LastName,
		Birth:     birthDate,
		Gender:    registerDto.Gender,
		Interests: registerDto.Interests,
		City:      registerDto.City,
		Password:  string(hashPassword),
		Email:     strings.ToLower(registerDto.Email),
	}, nil
}

func (u *UserEntity) GetID() uuid.UUID {
	return u.ID
}

func (u *UserEntity) GetFirstName() string {
	return u.FirstName
}

func (u *UserEntity) GetLastName() string {
	return u.LastName
}
func (u *UserEntity) GetBirth() time.Time {
	return u.Birth
}
func (u *UserEntity) GetGender() string {
	return u.Gender
}
func (u *UserEntity) GetInterest() string {
	return u.Interests
}
func (u *UserEntity) GetCity() string {
	return u.City
}
func (u *UserEntity) GetEmail() string {
	return u.City
}
func (u *UserEntity) GetPassword() string {
	return u.Password
}

func (u *UserEntity) ValidatePassword(inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword))
	return err
}
