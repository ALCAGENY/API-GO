package application

import (
	"api-go/src/product/domain/entities"
	"api-go/src/product/domain/ports"
)

type CreateProductUseCase struct {
	ProductRepository ports.ProductRepository
}

func NewCreateProductUseCase(productRepository ports.ProductRepository) *CreateProductUseCase{
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (p *CreateProductUseCase) Run(Name, Fecha_Adquisicion string) (entities.Product, error) {
	product := entities.Product{
		Name: Name,
		Fecha_Adquisicion: Fecha_Adquisicion,
	}

	newProduct, err := p.ProductRepository.Create(product)

	if err != nil {
		return entities.Product{}, err
	}

	return newProduct, nil
}