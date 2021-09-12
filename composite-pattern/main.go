package main

import (
	"design-pattern/composite-pattern/service/impl"
)

func main() {
	file1 := &impl.File{Name: "File1"}
	file2 := &impl.File{Name: "File2"}
	file3 := &impl.File{Name: "File3"}

	folder1 := &impl.Folder{Name: "Folder1"}

	folder1.Add(file1)

	folder2 := &impl.Folder{Name: "Folder2"}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
