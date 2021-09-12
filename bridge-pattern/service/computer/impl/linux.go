package impl

import (
	"fmt"

	"design-pattern/bridge-pattern/service/printer"
)

type Ubuntu struct {
	printer printer.Printer
}

func (u *Ubuntu) Print() {
	fmt.Println("Print Request from Ubuntu")
	u.printer.PrintFile()
}

func (u *Ubuntu) SetPrinter(printer printer.Printer) {
	u.printer = printer
}
