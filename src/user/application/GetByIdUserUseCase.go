package application

import (
	"api-go/src/user/domain/entities"
	"api-go/src/user/domain/ports"
)

type GetByIdUserUseCase struct {
	UserRepository ports.IUserRepository
}

func NewGetByIdUserUseCase(userRepository ports.IUserRepository) *GetByIdUserUseCase{
	return &GetByIdUserUseCase{ UserRepository: userRepository }
}

func (u *GetByIdUserUseCase) Run(id int64) (entities.User, error){
	user, err := u.UserRepository.GetByID(id)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}