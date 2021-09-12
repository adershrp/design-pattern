package computer

import (
	"design-pattern/bridge-pattern/service/printer"
)

type Computer interface {
	Print()
	SetPrinter(printer printer.Printer)
}
