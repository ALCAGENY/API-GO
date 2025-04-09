package controllers

import (
	"api-go/src/product/application"
	responses "api-go/src/shared/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductByIdController struct {
	GetByIdUseCase *application.GetProductByIdUseCase
}

func NewGetProductByIdController(getByIdUseCase *application.GetProductByIdUseCase) *GetProductByIdController {
	return &GetProductByIdController{GetByIdUseCase: getByIdUseCase}
}

func (c *GetProductByIdController) Run(ctx *gin.Context){
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

	product, err := c.GetByIdUseCase.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error getting product",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Product retrieved successfully",
		Data: product,
		Error: nil,
	})
}