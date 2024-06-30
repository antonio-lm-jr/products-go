package usecases

import (
	"errors"

	"github.com/antonio-lm-jr/products-go/application/products/domain/entity"
	"github.com/antonio-lm-jr/products-go/application/products/domain/valueobject"
	"github.com/antonio-lm-jr/products-go/application/products/dtos"
	"github.com/antonio-lm-jr/products-go/application/products/repository"
)

var (
	ErrorProductNotFound = errors.New("Product not found")
)

type ProductUsecaseImpl struct {
	repository repository.ProductRepositoryInterface
}

func ProvideProductUsecase(r repository.ProductRepositoryInterface) *ProductUsecaseImpl {
	return &ProductUsecaseImpl{
		repository: r,
	}
}

func (uc *ProductUsecaseImpl) ExecuteCreateProduct(input dtos.ProductInput) (dtos.ProductOutput, error) {
	code := valueobject.CreateIdentifier()

	name, err := valueobject.CreateName(input.Name)

	if err != nil {
		return dtos.ProductOutput{}, err
	}

	description, err := valueobject.CreateDescription(input.Description)

	if err != nil {
		return dtos.ProductOutput{}, err
	}

	price, err := valueobject.CreatePrice(input.Price)

	if err != nil {
		return dtos.ProductOutput{}, err
	}

	product := entity.Create(
		code.Value(),
		name.Value(),
		description.Value(),
		price.Value(),
	)

	uc.repository.CreateProduct(product)

	return dtos.ProductOutput{
		Identifier:  product.Identifier,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (uc *ProductUsecaseImpl) ExecuteGetProduct(identifier string) (dtos.ProductOutput, error) {
	product := uc.repository.GetProduct(identifier)

	if product.Identifier == "" {
		return dtos.ProductOutput{}, ErrorProductNotFound
	}

	return dtos.ProductOutput{
		Identifier:  product.Identifier,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (uc *ProductUsecaseImpl) ExecuteDeleteProduct(identifier string) {
	uc.repository.DeleteProduct(identifier)
}
