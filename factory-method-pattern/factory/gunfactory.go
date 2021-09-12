package factory

import (
	"fmt"

	"design-pattern/factory-method-pattern/product"
	"design-pattern/factory-method-pattern/product/impl"
)

func GetGun(gunType string) (product.IGun, error) {
	switch gunType {
	case "ak47":
		return impl.NewAk47(), nil

	case "musket":
		return impl.NewMusket(), nil

	}
	return nil, fmt.Errorf("wrong gun type")
}
