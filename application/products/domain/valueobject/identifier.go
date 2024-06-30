package valueobject

import (
	"github.com/segmentio/ksuid"
)

type Identifier struct {
	value string
}

func (i *Identifier) Value() string {
	return i.value
}

func CreateIdentifier() Identifier {
	return Identifier{
		value: ksuid.New().String(),
	}
}
