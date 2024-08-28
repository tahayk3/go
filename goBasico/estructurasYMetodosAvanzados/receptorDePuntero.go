package main

import "fmt"

type Rectangulo struct {
	ancho int
	alto  int
}

// Metodo con receptor de puntero
func (r *Rectangulo) Escalar(factor int) {
	r.ancho *= factor
	r.alto *= factor
}

/*
Un receptor de puntero recibe un puntero a la estructura original. Cualquier modificación dentro del
método afectará la estructura original.
*/
func main() {
	rectangulo1 := Rectangulo{ancho: 10, alto: 5}
	rectangulo1.Escalar(2)
	fmt.Println("Dimensiones escaladas:", rectangulo1.ancho, rectangulo1.alto) // 20, 10

}
