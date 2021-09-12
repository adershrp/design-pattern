package main

import (
	computer "design-pattern/bridge-pattern/service/computer/impl"
	printer "design-pattern/bridge-pattern/service/printer/impl"
)

func main() {

	inkJet := &printer.InkJet{}
	epson := &printer.Epson{}

	mac := &computer.Mac{}
	windows := &computer.Windows{}
	ubuntu := &computer.Ubuntu{}

	mac.SetPrinter(inkJet)
	mac.Print()

	mac.SetPrinter(epson)
	mac.Print()

	windows.SetPrinter(epson)
	windows.Print()

	windows.SetPrinter(inkJet)
	windows.Print()

	ubuntu.SetPrinter(inkJet)
	ubuntu.Print()

}
