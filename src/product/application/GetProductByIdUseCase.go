package application

import (
	"api-go/src/product/domain/entities"
	"api-go/src/product/domain/ports"
)

type GetProductByIdUseCase struct {
	ProductRepository ports.ProductRepository
}

func NewGetProductByIdUseCase(productRepository ports.ProductRepository) *GetProductByIdUseCase{
	return &GetProductByIdUseCase{ProductRepository: productRepository}
}

func (p *GetProductByIdUseCase) Run(id int64) (entities.Product, error){
	product, err := p.ProductRepository.GetByID(id)

	if err != nil {
		return entities.Product{}, err
	}

	return product, nil
}