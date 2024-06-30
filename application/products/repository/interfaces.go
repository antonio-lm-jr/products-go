package repository

import "github.com/antonio-lm-jr/products-go/application/products/domain/entity"

type ProductRepositoryInterface interface {
	CreateProduct(input entity.Product)
	GetProduct(identifier string) entity.Product
	DeleteProduct(identifier string)
}
