package controllers

import (
	"api-go/src/product/application"
	responses "api-go/src/shared/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	DeleteUseCase *application.DeleteProductUseCase
}

func NewDeleteProductController(deleteUseCase *application.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{DeleteUseCase: deleteUseCase}
}

func (d *DeleteProductController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil{
		ctx.JSON(http.StatusBadGateway, responses.Response{
			Success: false,
			Message: "Error parsing id",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	deleted, err := d.DeleteUseCase.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error deleting product",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Product deleted successfully",
		Data: deleted,
		Error: nil,
	})
}