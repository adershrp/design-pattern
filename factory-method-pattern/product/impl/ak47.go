package impl

import (
	"design-pattern/factory-method-pattern/product"
)

type ak47 struct {
	gun
}

func NewAk47() product.IGun {
	return &ak47{gun: gun{
		name: "AK 47", power: 89},
	}
}
