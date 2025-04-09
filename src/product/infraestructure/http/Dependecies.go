package http

import (
	"api-go/src/product/application"
	"api-go/src/product/domain/ports"
	"api-go/src/product/infraestructure/adapters"
	"api-go/src/product/infraestructure/http/controllers"
	"log"
)

var productRepository ports.ProductRepository

func init() {
	var err error

	productRepository, err = adapters.NewProductRepositoryMysql()

	if err != nil {
		log.Fatalf("Error creating user repository: %v", err)
	}
}

func SetUpCreateController() *controllers.CreateProductController {
	createUseCase := application.NewCreateProductUseCase(productRepository)
	return controllers.NewCreateProductController(createUseCase)
}

func SetUpGetIDController() *controllers.GetProductByIdController {
	getByIdUseCase := application.NewGetProductByIdUseCase(productRepository)
	return controllers.NewGetProductByIdController(getByIdUseCase)
}

func SetUpDeleteController() *controllers.DeleteProductController {
	deleteUseCase := application.NewDeleteProductUseCase(productRepository)
	return controllers.NewDeleteProductController(deleteUseCase)
}

