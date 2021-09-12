package impl

import (
	"fmt"

	"design-pattern/bridge-pattern/service/printer"
)

type Mac struct {
	printer printer.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print Request for Mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer printer.Printer) {
	m.printer = printer
}
