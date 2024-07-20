package main

import "fmt"

func main() {

	/*	Es como un arrary pero para acceder a los datos, ya se utiliza una posicion en el array, se accede por
		medio de una clave, por lo que la estructura de un mapa en go es de tipo: clave-valor.

		Go no garantiza el orden en los mapas, en base a cuando se ingresan datos.
	*/

	var mapa1 = map[string]string{
		"Venezuela": "Caracas",
		"Chile":     "Santiango",
	}

	//Ingresando a un elemento por medio de la clave
	fmt.Println(mapa1["Chile"])

	//Ingresando un nuevo elemento
	mapa1["Guatemala"] = "Guatemala"

	//borrando datos por medio de clave
	delete(mapa1, "Guatemala")

	//salida de datos
	fmt.Println(mapa1)

}
