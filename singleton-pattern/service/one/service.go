package one

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type oneService struct {
	name string
}

func (s *oneService) PrintSomething() {
	fmt.Println(s.name)
}

var lock = &sync.Mutex{}

var instance *oneService

func GetInstance() *oneService {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &oneService{strconv.Itoa(rand.Int())}
			fmt.Println("Creating new Instance of OneService")
		}
	}
	return instance
}
