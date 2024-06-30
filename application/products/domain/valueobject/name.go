package valueobject

import "errors"

type Name struct {
	value string
}

func (i *Name) Value() string {
	return i.value
}

func CreateName(input string) (Name, error) {
	const (
		_min = 3
		_max = 50
	)

	if len(input) < _min {
		return Name{}, errors.New("Name must have more than 3 characters")
	}

	if len(input) > _max {
		return Name{}, errors.New("Name must have less than 50 characters")
	}

	return Name{value: input}, nil
}
