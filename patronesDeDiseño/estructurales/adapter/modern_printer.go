package main

import "fmt"

//modernPrinter es la impresora moderna que imprime archivos pdf
type ModernPrinter struct{}

func (p *ModernPrinter) PrintFile() {
	fmt.Println("Imprimiendo archivo en formato PDF")
}
