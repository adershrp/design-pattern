package impl

import (
	"design-pattern/factory-method-pattern/product"
)

type musket struct {
	gun
}

func NewMusket() product.IGun {
	return &musket{
		gun: gun{name: "musket", power: 87},
	}
}
