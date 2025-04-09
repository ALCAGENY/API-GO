package application

import "api-go/src/product/domain/ports"

type DeleteProductUseCase struct {
	ProductRepository ports.ProductRepository
}

func NewDeleteProductUseCase(productRepository ports.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{ProductRepository: productRepository}
}

func (p *DeleteProductUseCase) Run(id int64) (bool, error){
	deleted, err := p.ProductRepository.Delete(id)

	if err != nil {
		return false, err
	}

	return deleted, nil
}