package impl

import (
	"fmt"

	"design-pattern/composite-pattern/service"
)

type Folder struct {
	components []service.Component
	Name       string
}

func (f *Folder) GetName() string {
	return f.Name
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c service.Component) {
	f.components = append(f.components, c)
}
