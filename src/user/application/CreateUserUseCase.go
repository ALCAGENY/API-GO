package application

import (
	"api-go/src/user/application/services"
	"api-go/src/user/domain/entities"
	"api-go/src/user/domain/ports"
)


type CreateUserUseCase struct {
	UserRepository ports.IUserRepository
	EncrtyptService services.EncrtyptService
}

func NewCreateUserUseCase(userRepository ports.IUserRepository, encryptService services.EncrtyptService) *CreateUserUseCase{
	return &CreateUserUseCase{
		UserRepository: userRepository,
		EncrtyptService: encryptService,
	}
}


func (u *CreateUserUseCase) Run(Name, Email, Password string) (entities.User, error){
	
	hashpwd, err := u.EncrtyptService.Encrypt([]byte(Password))

	if err != nil {
		return entities.User{}, err
	}
	
	user := entities.User{
		Name:     Name,
		Email:    Email,
		Password: hashpwd,
	}

	newUser, err := u.UserRepository.Create(user)
	
	if err != nil {
		return entities.User{}, err
	}
	
	return newUser, nil

}