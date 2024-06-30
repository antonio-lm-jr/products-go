package controllers

import (
	"github.com/antonio-lm-jr/products-go/application/products/dtos"
	"github.com/antonio-lm-jr/products-go/application/products/usecases"
	"github.com/gofiber/fiber/v2"
)

type ProductControllerImpl struct {
	usecase usecases.ProductUsecaseInterface
}

func ProvideProductController(u usecases.ProductUsecaseInterface) *ProductControllerImpl {
	return &ProductControllerImpl{
		usecase: u,
	}
}

func (pc *ProductControllerImpl) CreateProductController(ctx *fiber.Ctx) error {
	productInput := new(dtos.ProductInput)

	if err := ctx.BodyParser(productInput); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	actionResult, err := pc.usecase.ExecuteCreateProduct(*productInput)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(actionResult)
}

func (pc *ProductControllerImpl) GetProductController(ctx *fiber.Ctx) error {
	identifier := ctx.Params("id")

	actionResult, err := pc.usecase.ExecuteGetProduct(identifier)

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(
		actionResult,
	)
}

func (pc *ProductControllerImpl) DeleteProductController(ctx *fiber.Ctx) error {
	identifier := ctx.Params("id")

	pc.usecase.ExecuteDeleteProduct(identifier)

	return ctx.JSON(fiber.Map{
		"message": "Product deleted",
	})
}
