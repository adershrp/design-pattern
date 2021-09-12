package factory

import (
	"design-pattern/abstract-factory-method-pattern/product"
	"design-pattern/abstract-factory-method-pattern/product/impl/shoe"
	"design-pattern/abstract-factory-method-pattern/product/impl/short"
)

type nike struct {
}

func (n *nike) MakeShoe() product.IShoe {
	return &shoe.NikeShoe{Shoe: shoe.Shoe{Logo: "Nike", Size: 14}}

}

func (n *nike) MakeShort() product.IShort {
	return &short.NikeShort{Short: short.Short{Logo: "Nike", Size: 10}}

}
