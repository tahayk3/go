package main

import "fmt"

/*
Un puntero es una variable que almacena la dirección de memoria de otra variable. En lugar de almacenar
 directamente un valor, un puntero almacena la ubicación en la memoria donde ese valor se encuentra.
*/

/*
	¿Por qué usar punteros?
	Eficiencia: En lugar de copiar grandes estructuras o arrays, puedes pasar un puntero, que es simplemente una
	dirección de memoria, lo que es mucho más ligero.

	Modificar el valor original: Si pasas una variable por valor a una función, cualquier modificación de esa
	variable dentro de la función no afectará a la variable original. Sin embargo, si pasas un puntero, puedes
	modificar el valor original.
*/
func main() {
	//puntero de tipo int, creado pero sin apuntar a una variable var p *int
	//asignacion de puntero
	var num1 int = 10
	var p *int = &num1
	fmt.Println("La direccion de memoria de la variable num1 es ", p)

	/*	Para acceder al valor almacenado en la direccion de la memoria a la que apunta un
		puntero, se usa * dereferenciación.
	*/
	fmt.Println(*p)

	//modificar el valor por medio del puntero
	*p = 20
	fmt.Println("nuevo valor modificando por medio del puntero", num1)

}
