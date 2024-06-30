package entity

type Product struct {
	Identifier  string
	Name        string
	Description string
	Price       float64
}

func Create(identifier string, name string, description string, price float64) Product {
	return Product{
		Identifier:  identifier,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
