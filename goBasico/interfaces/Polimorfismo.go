package main

import "fmt"

type Describible interface {
	Describir() string
}

type Persona struct {
	nombre string
	edad   int
}

func (p Persona) Describir() string {
	return fmt.Sprintf("%s tiene %d años", p.nombre, p.edad)

}

type Producto struct {
	nombre string
	precio float64
}

func (p Producto) Describir() string {
	return fmt.Sprintf("El producto %s tiene un precio de %.2f", p.nombre, p.precio)
}

// Función que recibe cualquier tipo que implemente la interfaz Describible
func imprimirDescripcion(d Describible) {
	fmt.Println(d.Describir())
}

func main() {
	persona1 := Persona{nombre: "Juan", edad: 20}
	producto1 := Producto{nombre: "taza", precio: 15}

	//se esta creando polimorfisno, dependiendo del objeto, la funcion, imprimirDescripcion modifica su comportamiento
	imprimirDescripcion(persona1)
	imprimirDescripcion(producto1)

}
