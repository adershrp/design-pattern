package main

import (
	"design-pattern/singleton-pattern/service/one"
	"design-pattern/singleton-pattern/service/three"
	"design-pattern/singleton-pattern/service/two"
)

func main() {
	for i := 0; i < 10; i++ {

		oneService := one.GetInstance()
		oneService.PrintSomething()

		twoService := two.GetInstance()
		twoService.DoSomething()

		threeService := three.GetInstance()
		threeService.DoSomething()
	}
}
