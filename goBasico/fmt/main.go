package main

import "fmt"

/*	Se utiliza para mostrar texto en terminal, escribir string(asignarlos a variables), para manejo
	de archivos de texto y para enviar texto como respuesta a request de un servidor.
*/

//fmt puede recibir una cantidad variable de parametros

var a int
var b string = "Programa"
var c bool
var d bool = true

func main() {
	e := 42
	f := "String interpretado"
	g := `String no    interpretado`

	//cantidad variable de parametros
	fmt.Println(e, f, g)

	//printf formatea en base al especificador de formato(v% es el valor por defecto)
	fmt.Printf("El valor de a es: %v\n", a)
	fmt.Printf("El tipo de a es: %T\n", a)
	fmt.Printf("El valor de b es: %v\n", b)
	fmt.Printf("El tipo de b es: %T\n", b)

	//asignar string a una variable con fmt
	frase := fmt.Sprint("La frase es ", f, " jiji")
	fmt.Println(frase)
}
