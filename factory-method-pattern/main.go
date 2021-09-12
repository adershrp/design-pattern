package main

import (
	"fmt"

	"design-pattern/factory-method-pattern/factory"
	"design-pattern/factory-method-pattern/product"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g product.IGun) {
	fmt.Printf("gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("power: %d", g.GetPower())
	fmt.Println()
}
