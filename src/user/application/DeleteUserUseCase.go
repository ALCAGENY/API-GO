package application

import (
	"api-go/src/user/domain/ports"
)

type DeleteUserUseCase struct {
	UserRepository ports.IUserRepository
}


func NewDeleteUserUseCase(userRepository ports.IUserRepository)*DeleteUserUseCase{
	return &DeleteUserUseCase{UserRepository: userRepository}
}

func(u *DeleteUserUseCase)Run(id int64) (bool, error){
	deleted, err := u.UserRepository.Delete(id)

	if err != nil {
		return false, err
	}

	return deleted, nil
}