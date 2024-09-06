/*
	Para este ejercicio

	Crea tu propio tipo. Haz que el tipo subyacente, raíz o implícito sea un int.

	Crea una VARIABLE de tu nuevo TIPO con el IDENTIFICADOR “x” usando la palabra clave “VAR”

	En func main

	Imprime el valor de la variable “x”

	Imprime el tipo de la variable “x”

	Asigna 42 a la VARIABLE “x” usando el OPERADOR “=”

	Imprime el valor de la variable “x”
*/

package main

import "fmt"

type variablePropia int

var x variablePropia

func main() {
	fmt.Println(x)
	fmt.Printf("El tipo es: %T", x)
	x = 42
	fmt.Println(x)

}
