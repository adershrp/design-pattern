package impl

import (
	"fmt"
)

type Epson struct {
}

func (e Epson) PrintFile() {
	fmt.Println("Printing by Epson Printer")
}
