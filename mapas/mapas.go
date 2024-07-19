package main

import "fmt"

func main() {

	/*	Es como un arrary pero para acceder a los datos, ya se utiliza una posicion en el array, se accede por
		medio de una clave, por lo que la estructura de un mapa en go es de tipo: clave-valor
	*/
	var mapa1 = map[string]string{
		"Venezuela": "Caracas",
		"Chile":     "Santiango",
	}
	fmt.Println(mapa1)

}
