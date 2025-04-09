package http

import (
	"api-go/src/user/application"
	"api-go/src/user/application/services"
	"api-go/src/user/domain/ports"
	"api-go/src/user/infraestructure/adapters"
	"api-go/src/user/infraestructure/http/controllers"
	"api-go/src/user/infraestructure/http/controllers/helper"
	"log"
)

var (
	userRepository  ports.IUserRepository
	encrypyService services.EncrtyptService
)

func init() {
	var err error

	userRepository, err = adapters.NewUserRepository()
	if err != nil {
		log.Fatalf("Error creating user repository: %v", err)
	}

	encrypyService, err = helper.NewEncryptHelper()
	if err != nil {
		log.Fatalf("Error creating encrypt service: %v", err)
	}
}

func SetUpCreateController() *controllers.CreateUserController{
	createUseCase := application.NewCreateUserUseCase(userRepository, encrypyService)
	return controllers.NewCreateUserController(createUseCase)
}

func SetUpAuthController() *controllers.AuthController{
	authUseCase := application.NewAuthUseCase(userRepository)
	return controllers.NewAuthController(authUseCase)
}

func SetUpGetIDController() *controllers.GetByIdController{
	getByIdUseCase := application.NewGetByIdUserUseCase(userRepository)
	return controllers.NewGetByIdController(getByIdUseCase)
}

func SetUpDeleteController() *controllers.DeleteUserController{
	deleteUserUseCase := application.NewDeleteUserUseCase(userRepository)
	return controllers.NewDeleteUserController(deleteUserUseCase)
}