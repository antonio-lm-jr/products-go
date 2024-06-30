package main

import (
	"github.com/antonio-lm-jr/products-go/application/products/usecases"
	"github.com/antonio-lm-jr/products-go/config"
	"github.com/antonio-lm-jr/products-go/controllers"
	"github.com/antonio-lm-jr/products-go/http"
	"github.com/antonio-lm-jr/products-go/infrastructure/persistence"
	"github.com/antonio-lm-jr/products-go/infrastructure/repository"
)

func initHttpService() http.HttpServiceInterface {
	config := config.ProvideConfig()
	database := persistence.ProvideDatabase(config)

	productRepoImpl := repository.ProvideProductRepository(database)

	productUsecaseImpl := usecases.ProvideProductUsecase(productRepoImpl)
	ProductControllerImpl := controllers.ProvideProductController(productUsecaseImpl)

	return http.ProvideFiberServer(
		ProductControllerImpl,
		config,
	)
}

func main() {
	initHttpService().Init()
}
