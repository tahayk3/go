/*
Ejercicio 2: Sumar Números Concurrentemente
Crea un programa que tome un array de enteros y lance dos goroutines: una que calcule la suma de la primera mitad del array y otra
que calcule la suma de la segunda mitad. Luego, usa un canal para recoger las dos sumas y muestra el resultado total.
*/

package main

import (
	"fmt"
)

func firstFunction(arrayNumeros [10]int, canal chan int) {
	acum := 0
	for x := 0; x <= 4; x++ {
		acum = acum + arrayNumeros[x]
	}
	canal <- acum
}

func secondFunction(arrayNumeros [10]int, canal chan int) {
	acum := 0
	for x := 9; x >= 5; x-- {
		acum = acum + arrayNumeros[x]
	}
	canal <- acum
}

func main() {
	arrayNumeros := [10]int{5, 2, 15, 50, 1000, 3, 5, 9, 13, 2}
	canal := make(chan int)

	go firstFunction(arrayNumeros, canal)
	go secondFunction(arrayNumeros, canal)

	suma1 := <-canal
	suma2 := <-canal

	fmt.Println("La primera suma es: ", suma1)
	fmt.Println("La segunda suma es: ", suma2)

	sumaTotal := suma1 + suma2
	fmt.Println(sumaTotal)

	fmt.Println("Fin del programa")
}
