package main

/*
Imagina que estás desarrollando un sistema que necesita comunicarse con diferentes tipos de impresoras. Tienes una impresora antigua que
solo puede recibir comandos en formato PCL (Printer Command Language) y una impresora moderna que usa comandos en formato PDF. El sistema
necesita enviar documentos a ambas impresoras sin modificar su lógica interna. Aquí es donde entra el patrón Adapter.
*/

//correr con go mod init adapter y go run .
func main() {
	modernPrinter := &ModernPrinter{}
	oldPrinter := &OldPrinter{}
	adapter := &OldPrinterAdapter{oldPrinter: oldPrinter}

	// Usamos la impresora moderna directamente
	modernPrinter.PrintFile()

	// Usamos la impresora antigua a través del adaptador
	adapter.PrintFile()
}
