package main

import (
	"fmt"
	"godesignpattern/structural_pattern/bridge"
)

func main() {
	hpPrinter := &bridge.Hp{}
	epsonPrinter := &bridge.Epson{}

	macComputer := &bridge.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &bridge.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
