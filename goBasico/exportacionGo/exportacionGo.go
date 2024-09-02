package main

import "fmt"

/*
En Go, el concepto de "métodos accesores" como en otros lenguajes (como public, private, protected) se maneja de
manera diferente. Go no utiliza palabras clave específicas para la visibilidad de métodos o propiedades
como en otros lenguajes orientados a objetos. En su lugar, Go utiliza la convención de
nombres basándose en la capitalización de las letras.
*/

/*
Reglas en go para exportar:
Exportado (Público): Si el nombre de una función, método, estructura o campo de estructura comienza
 con una letra mayúscula, se exporta y es accesible desde otros paquetes. Esto es similar a public en otros
  lenguajes.

No exportado (Privado): Si el nombre comienza con una letra minúscula, no se exporta y solo es
accesible dentro del mismo paquete. Esto es equivalente a private en otros lenguajes.
*/

type Persona struct {
	//exportado, lo pueden ver otros paquetes
	Nombre string
	//no exportado, no lo pueden ver otros paquetes
	edad int
}

//Exportado
func (p Persona) Saludar() {
	fmt.Println("Hola, mi nombre es %s\n", p.Nombre)
}

//No exportado
func (p Persona) calcularEdad() int {
	return p.edad
}

func main() {
	p := Persona{Nombre: "Juanita", edad: 1}

	p.Saludar()           // Funciona
	fmt.Println(p.Nombre) // Funciona
	fmt.Println(p.edad)   // No funciona, edad es privado
	p.calcularEdad()      // No funciona, calcularEdad es privado
}
