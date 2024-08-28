package main

import "fmt"

type Rectangulo struct {
	ancho int
	alto  int
}

// Metodo con recepetor de valor
func (r Rectangulo) Area() int {
	return r.ancho * r.alto
}

/*
	Un receptor de valor recibe una copia de la estructura. Cualquier modificación dentro del método no
	afectará la estructura original.
*/
func main() {
	rectangulo1 := Rectangulo{ancho: 15, alto: 5}
	fmt.Println("Area:", rectangulo1.Area())
}
