package repository

import (
	"github.com/antonio-lm-jr/products-go/application/products/domain/entity"
	"github.com/antonio-lm-jr/products-go/infrastructure/persistence"
	"log"
)

type PersonRepositoryImpl struct {
	Database *persistence.Database
}

func ProvideProductRepository(database *persistence.Database) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{Database: database}
}

func (pr *PersonRepositoryImpl) CreateProduct(input entity.Product) {
	prepareStatement, err := pr.Database.Conn.Prepare(
		"insert into products (identifier, name, description, price) values (?, ?, ?, ?);",
	)

	if err != nil {
		log.Println("Error on prepare statement: ", err.Error())
	}

	_, err = prepareStatement.Exec(input.Identifier, input.Name, input.Description, input.Price)

	if err != nil {
		log.Println("Error on execute statement: ", err.Error())
	}
}

func (pr *PersonRepositoryImpl) GetProduct(identifier string) entity.Product {
	product := entity.Product{}

	row := pr.Database.Conn.QueryRow(
		"select identifier, name, description, price from products where identifier = ?;", identifier,
	)

	err := row.Scan(
		&product.Identifier, &product.Name, &product.Description, &product.Price,
	)

	if err != nil {
		log.Println("Error on scan row: ", err.Error())
	}

	return product
}

func (pr *PersonRepositoryImpl) DeleteProduct(identifier string) {
	prepareStatement, err := pr.Database.Conn.Prepare(
		"delete from products where identifier = ?;",
	)

	if err != nil {
		log.Println("Error on prepare statement: ", err.Error())
	}

	_, err = prepareStatement.Exec(identifier)

	if err != nil {
		log.Println("Error on execute statement: ", err.Error())
	}
}
