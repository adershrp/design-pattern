package two

import (
	"fmt"
	"math/rand"
	"sync"
)

type twoService struct {
	random int
}

func (s *twoService) DoSomething() {
	fmt.Println(s.random)
}

var instance *twoService
var once sync.Once

func GetInstance() *twoService {
	if instance == nil {
		once.Do(func() {
			instance = &twoService{random: rand.Int()}
		})
	}
	return instance
}
