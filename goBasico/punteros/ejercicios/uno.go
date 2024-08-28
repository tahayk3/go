/*
Ejercicio: Crea una función en Go llamada incrementar que reciba un puntero a un entero y aumente
su valor en 1. Luego, prueba esta función en main con una variable entera.
*/
package main

import "fmt"

func incrementar(incremento *int) {
	paso := *incremento
	*incremento = paso + 1
}
func main() {
	var num1 int = 0
	var incremento *int = &num1
	incrementar(incremento)
	incrementar(incremento)
	incrementar(incremento)
	incrementar(incremento)
	incrementar(incremento)
	incrementar(incremento)
	fmt.Println("El valor del entero es: ", num1)
}
