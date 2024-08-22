package main

import "fmt"

//oldPrinter es la impresora antigua que imprime archivos PCL
type OldPrinter struct{}

func (p *OldPrinter) PrinterPCLFile() {
	fmt.Println("Imprimiendo archivo en formato PCL")
}
