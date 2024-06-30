package usecases

import (
	"github.com/antonio-lm-jr/products-go/application/products/dtos"
)

type ProductUsecaseInterface interface {
	ExecuteCreateProduct(input dtos.ProductInput) (dtos.ProductOutput, error)
	ExecuteGetProduct(identifier string) (dtos.ProductOutput, error)
	ExecuteDeleteProduct(identifier string)
}
