package controllers

import (
	"api-go/src/shared/response"
	"api-go/src/user/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	DeleteUseCase *application.DeleteUserUseCase
}

func NewDeleteUserController(deleteUseCase *application.DeleteUserUseCase) *DeleteUserController{
	return&DeleteUserController{DeleteUseCase: deleteUseCase}
}

func (ctrl *DeleteUserController)Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error parsing id",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	deleted, err := ctrl.DeleteUseCase.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error deleting user",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "User deleted",
		Data: deleted,
		Error: nil,
	})
}