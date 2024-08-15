/*
Ejercicio 1: Imprimir Números Concurrentemente
Crea un programa que lance dos goroutines, una que imprima los números pares del 0 al 10 y otra que imprima los números impares del 1 al 9.
Haz que ambas goroutines se ejecuten de manera concurrente.
*/

package main

import (
	"fmt"
	"time"
)

func numerosPares() {
	for x := 0; x <= 10; x += 2 {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func numerosImpares() {
	for x := 1; x <= 9; x += 2 {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go numerosPares()
	go numerosImpares()
	time.Sleep(10 * time.Second)
}
