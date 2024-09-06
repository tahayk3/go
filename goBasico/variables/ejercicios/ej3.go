/*
	Usando el codigo del ejercio anterior,
	1. En scope a nivel de paquete, asigna los siguientes valores a las tres variables
		a. a x asignarle 42
		b. a y asignale "James Bond"
		c. a z asignale true
	2. en main
		a. Usa fmt.Sprintf para imprimir todos los VALORES en un solo string. ASIGNA el valor
		   retornado de TIPO string usando el operador de declaracion corta a la VARIABLE con el
		   IDENTIFICADOR "s"
		b. Imprime el valor almacenado por la variables "s"

*/

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(s)

}
