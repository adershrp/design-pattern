package impl

import (
	"fmt"

	"design-pattern/bridge-pattern/service/printer"
)

type Windows struct {
	printer printer.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request from Windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(printer printer.Printer) {
	w.printer = printer
}
