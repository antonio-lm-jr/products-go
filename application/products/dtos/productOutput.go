package dtos

type ProductOutput struct {
	Identifier  string  `json:"identifier"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
