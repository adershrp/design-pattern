package impl

import (
	"fmt"
)

type File struct {
	Name string
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}
