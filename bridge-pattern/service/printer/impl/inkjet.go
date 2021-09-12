package impl

import (
	"fmt"
)

type InkJet struct {
}

func (i InkJet) PrintFile() {
	fmt.Println("Printing by InkJet printer")
}
