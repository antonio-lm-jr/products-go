package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type ProductControllerInterface interface {
	CreateProductController(ctx *fiber.Ctx) error
	GetProductController(ctx *fiber.Ctx) error
	DeleteProductController(ctx *fiber.Ctx) error
}
