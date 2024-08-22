// Las interfaces permiten escribir código más flexible y reutilizable.
// Una interface en Go es un tipo que define un conjunto de métodos, pero no los implementa.

package main

import "fmt"

//definicion de la interfaz
type Describible interface {
	Describir() string
}

//Estructura que implementa la interfez
type Persona struct {
	nombre string
	edad   int
}

//Implementacion del metodo Describir
/* CT
De esta manera asociamos a un objeto de tipo Persona creado a partir de un tipo con la funcion Describir(), de esa manera
cumplimos con la interfaz
*/
func (p Persona) Describir() string {
	return fmt.Sprintf("%s tiene %d años", p.nombre, p.edad)
}

func main() {
	p := Persona{nombre: "Juan", edad: 20}

	//Persona implementa Describible, se puede asignar p a una variable de tipo Describible
	var d Describible = p

	fmt.Println(d.Describir())

}
