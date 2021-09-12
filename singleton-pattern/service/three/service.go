package three

import (
	"fmt"
	"math/rand"
)

type threeService struct {
	random int
}

func (s *threeService) DoSomething() {
	fmt.Println(s.random)
}

var instance *threeService

func init() {
	instance = &threeService{random: rand.Int()}
}

func GetInstance() *threeService {
	return instance
}
