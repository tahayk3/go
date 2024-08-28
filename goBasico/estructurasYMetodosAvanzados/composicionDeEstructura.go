package main

import "fmt"

/*
Go no tiene herencia como otros lenguajes orientados a objetos, pero puedes usar la composición
para reutilizar código. Puedes incrustar una estructura dentro de otra y acceder a sus campos y métodos
directamente.
*/

type Persona struct {
	nombre string
	edad   int
}

type Empleado struct {
	// Incrustación de la estructura Persona
	Persona
	empresa string
}

func main() {
	//crear objeto con una estructura dentro de otra
	empleado := Empleado{
		Persona: Persona{nombre: "Juan", edad: 30},
		empresa: "xd",
	}
	fmt.Println(empleado.edad)
}
