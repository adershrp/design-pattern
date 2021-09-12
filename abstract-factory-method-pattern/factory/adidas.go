package factory

import (
	"design-pattern/abstract-factory-method-pattern/product"
	"design-pattern/abstract-factory-method-pattern/product/impl/shoe"
	"design-pattern/abstract-factory-method-pattern/product/impl/short"
)

type adidas struct {
}

func (a *adidas) MakeShoe() product.IShoe {
	return &shoe.AdidasShoe{Shoe: shoe.Shoe{Logo: "adidas", Size: 14}}

}

func (a *adidas) MakeShort() product.IShort {
	return &short.AdidasShort{Short: short.Short{Logo: "adidas", Size: 10}}
}
