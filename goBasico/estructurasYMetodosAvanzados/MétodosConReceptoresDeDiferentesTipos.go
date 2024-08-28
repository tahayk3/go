/*
	Puedes definir métodos con receptores de distintos tipos de estructuras. Esto es útil cuando
 	diferentes estructuras comparten funcionalidad.
*/

package main

import "fmt"

type Circulo struct {
	radio int
}

type Cuadrado struct {
	lado int
}

// Método para calcular el área de un círculo
func (c Circulo) Area() float64 {
	return 3.14 * float64(c.radio) * float64(c.radio)
}

// Metodo para calcular el area de un cuadrado
func (c Cuadrado) Area() int {
	return c.lado * c.lado
}

func main() {
	circulo := Circulo{radio: 15}
	cuadrado := Cuadrado{lado: 15}
	fmt.Println(circulo.Area())
	fmt.Println(cuadrado.Area())
}
