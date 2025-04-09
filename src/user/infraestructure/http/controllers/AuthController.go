package controllers

import (
	"api-go/src/shared/middlewares"
	"api-go/src/shared/response"
	"api-go/src/user/application"
	"api-go/src/user/infraestructure/http/controllers/helper"
	"api-go/src/user/infraestructure/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthUseCase *application.AuthUseCase
	EncryptService *helper.EncryptHelper
}

func NewAuthController(authUseCase *application.AuthUseCase)*AuthController{
	return &AuthController{
		AuthUseCase: authUseCase,
	}
}


func (ctrl *AuthController)Run(ctx *gin.Context){
	var req request.AuthRequest

	if err := ctx.BindJSON(&req); err != nil{ 
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Error: err.Error(),
			Data: nil,
		})
		return 
	}


	user,  err := ctrl.AuthUseCase.Run(req.Email)

	if err != nil {
		switch err.Error(){
		case "sql: no rows in result set":
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false,
				Message: "El email no existe",
                Error: err.Error(),
                Data: nil,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, responses.Response{
				Success: false,
				Message: "Error al iniciar sesión",
                Error: err.Error(),
                Data: nil,
			})
		}

		return
	}

	if err := ctrl.EncryptService.Compare(user.Password, []byte(req.Password)); err != nil{
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false,
			Message: "Contraseña incorrecta",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	token, err := middlewares.GenerateJWT(int64(user.ID), user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al generar el token",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Inicio de sesión exitoso",
		Error: nil,
		Data: map[string]interface{}{
			"Token": token,
		},
	})


}
