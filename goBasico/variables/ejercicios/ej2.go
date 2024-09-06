/*
Usa var para DECLARAR tres VARIABLES. Las variables deben tener scope a nivel de paquete.
No asignar VALORES a las variables. Usa los siguientes IDENTIFICADORES para la variables y
asegurate de que las variables son del los siguientes TIPOS (lo que quiere decir es que pueden
almacenar VALORES de ese TIPO).

a. identificador "x" tipo int
b. identificador "y" tipo string
c. identificador "z" tipo bool

en main:

	a. imprime los valores de cada indentificador
	b. El compilador asigna valores a las variables ¿Coomo son llamados esos valores? por defecto, 0
*/
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
