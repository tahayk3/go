package main

/*
	OldPrinterAdapter es el adaptardor que permite a una OldPrinter ser utilizada
	cona una printer moderna
*/
type OldPrinterAdapter struct {
	oldPrinter *OldPrinter
}

func (a *OldPrinterAdapter) PrintFile() {
	a.oldPrinter.PrinterPCLFile()
}
