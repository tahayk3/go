package main

import "fmt"

func main() {
	//Formas de declarar una variable en go(strings)
	var primeraForma string = "texto 1"
	var segundaForma = "texto 2"
	terceraForma := "texto 3"

	//salida de variables
	fmt.Println("Salida: " + primeraForma)
	fmt.Println("Salida: " + segundaForma)
	fmt.Println("Salida: " + terceraForma)

	//Variables de tipo numerico, int16,32,64, tambien tiene la capacidad de trabajar con numeros completos
	var num1 int16 = 5
	fmt.Println(num1)

	/*Arreglos, otra forma seria: frutas: = [3]string{"Pera", "Manzana", "Uva"}*/
	var frutas = [3]string{}
	frutas[0] = "Pera"
	fmt.Println(frutas[0])

	//Slices, son como un arreglo, pero sin un tamaño definido
	var frutas2 = []string{"pera"}
	//Agregando elementos a slice con append
	frutas2 = append(frutas2, "uva")
	fmt.Println(frutas2)
}
