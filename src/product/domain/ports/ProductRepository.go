package ports

import "api-go/src/product/domain/entities"

type ProductRepository interface {
	Create(product entities.Product) (entities.Product, error)
	GetByID(id int64) (entities.Product, error)
	Delete(id int64) (bool, error)
}