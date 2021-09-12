package factory

import (
	"fmt"

	"design-pattern/abstract-factory-method-pattern/product"
)

type ISportsFactory interface {
	MakeShoe() product.IShoe
	MakeShort() product.IShort
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
