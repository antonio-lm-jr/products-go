package valueobject

import "errors"

type Description struct {
	value string
}

func (i *Description) Value() string {
	return i.value
}

func CreateDescription(input string) (Description, error) {
	const (
		_min = 3
		_max = 100
	)

	if len(input) < _min {
		return Description{}, errors.New("Name must have more than 3 characters")
	}

	if len(input) > _max {
		return Description{}, errors.New("Name must have less than 50 characters")
	}

	return Description{value: input}, nil
}
