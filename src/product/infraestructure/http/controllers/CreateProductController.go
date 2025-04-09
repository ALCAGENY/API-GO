package controllers

import (
	"api-go/src/product/application"
	"api-go/src/product/infraestructure/http/request"
	responses "api-go/src/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateProductController struct {
	createUseCase *application.CreateProductUseCase
	Validate validator.Validate
}

func NewCreateProductController(createUseCase *application.CreateProductUseCase) *CreateProductController{
	return &CreateProductController{
		createUseCase: createUseCase,
		Validate: *validator.New(),
	}
}

func (c *CreateProductController) Run(ctx *gin.Context){
	var req request.ProductRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error binding request",
			Data: nil,
			Error: err.Error(),
		})
		
		return 
	}

	if err := c.Validate.Struct(req); err != nil{
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Validation error",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	product, err := c.createUseCase.Run(req.Name, req.Fecha_Adquisicion)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error creating product",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Product created successfully",
		Data: product,
		Error: nil,
	})


}