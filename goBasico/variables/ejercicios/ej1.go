/*
1.	Usando el operador de declaracion corta, ASIGNA los siguientes VALORES a VARIABLES
con los IDENTIFICADORES "x", "y" y "z"
	a. 42
	b. "James Bond"
	c. true

2.	Luego imprime los valores almacenados en esas variables usando
	a. Una sola declaracion de la funcion println
	b. Multiles declaraciones de println
*/

package main

import "fmt"

func main() {
	a := 42
	b := "James Bond"
	c := true
	fmt.Println(a, b, c)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
