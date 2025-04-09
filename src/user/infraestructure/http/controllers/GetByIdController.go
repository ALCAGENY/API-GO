package controllers

import (
	"api-go/src/shared/response"
	"api-go/src/user/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByIdController struct{
	GetByIdUseCase *application.GetByIdUserUseCase
}

func NewGetByIdController(getByIdUseCase *application.GetByIdUserUseCase) *GetByIdController{
	return &GetByIdController{GetByIdUseCase: getByIdUseCase}
}


func (ctrl *GetByIdController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error parsing id",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	user, err := ctrl.GetByIdUseCase.Run(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error getting user",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "User found",
		Data: user,
		Error: nil,
	})


}