package main

import (
	"fmt"
     
)

type Computer interface {
	Print()
	SetPrinter(Printer)
}
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}
func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

type Printer struct{
	PrintFile()
}
type Epson struct{

}
func (p *Epson)Printfile(){
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct{

}
func (p *Hp)PrintFile(){
	fmt.Println("Printing by a HP Printer")
}
