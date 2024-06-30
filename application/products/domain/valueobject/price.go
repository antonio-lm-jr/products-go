package valueobject

import "errors"

type Price struct {
	value float64
}

const (
	_minValue = 0.1
	_maxValue = 100000.0
)

func (i *Price) Value() float64 {
	return i.value
}

func CreatePrice(input float64) (Price, error) {
	if input < _minValue {
		return Price{}, errors.New("Price must be greater than 0")
	}

	if input > _maxValue {
		return Price{}, errors.New("Price must be less than 100.000")
	}

	return Price{value: input}, nil
}
