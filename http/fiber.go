package http

import (
	"github.com/antonio-lm-jr/products-go/config"
	"github.com/antonio-lm-jr/products-go/controllers"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	controllerProduct controllers.ProductControllerInterface
	config            config.Config
}

func ProvideFiberServer(productController *controllers.ProductControllerImpl, config config.Config) *FiberServer {
	return &FiberServer{
		controllerProduct: productController,
		config:            config,
	}

}

func (server FiberServer) Init() {
	app := fiber.New()

	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "ok",
		})
	})

	app.Post("/products", server.controllerProduct.CreateProductController)
	app.Get("/products/:id", server.controllerProduct.GetProductController)
	app.Delete("/products/:id", server.controllerProduct.DeleteProductController)

	app.Listen(":5001")
}
