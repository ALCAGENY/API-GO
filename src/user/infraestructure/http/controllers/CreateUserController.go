package controllers

import (
	"api-go/src/shared/response"
	"api-go/src/user/application"
	"api-go/src/user/infraestructure/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUserController struct {
	CreateUSeCase *application.CreateUserUseCase
	validate *validator.Validate
}


func NewCreateUserController(createUSeCase *application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{
		CreateUSeCase: createUSeCase,
		validate: validator.New(),
	}
}

func (ctrl *CreateUserController)Run(ctx *gin.Context){

	var req request.CreateUserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error binding request",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	if err := ctrl.validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Validation error",
			Data: nil,
			Error: err.Error(),
		})
		
		return
	}

	user, err := ctrl.CreateUSeCase.Run(req.Name, req.Email, req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error creating user",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "User created successfully",
		Data: user,
		Error: nil,
	})


}

