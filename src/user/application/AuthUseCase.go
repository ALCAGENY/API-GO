package application

import (
	"api-go/src/user/domain/entities"
	"api-go/src/user/domain/ports"
)

type AuthUseCase struct {
	UserRepository ports.IUserRepository
}

func NewAuthUseCase(userRepository ports.IUserRepository) *AuthUseCase{
	return &AuthUseCase{UserRepository: userRepository}
}

func (u *AuthUseCase)Run(email string) (entities.User, error){
	user, err := u.UserRepository.GetByEmail(email)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}