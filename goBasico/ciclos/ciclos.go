package main

import (
	"fmt"
	"strings"
)

func main() {
	//for
	for x := 0; x <= 100; x++ {
		fmt.Println("El numero es: ", x)
	}

	var mapa1 = map[string]string{
		"Venezuela": "Caracas",
		"Chile":     "Santiango",
	}

	//"foreach" utilizado para recorrer un mapa
	for c, v := range mapa1 {
		fmt.Println("La capital de " + c + " es: " + v)
	}

	//"mientras"
	var fruta string = ""
	for {
		if fruta == "sandia" {
			break
		} else {
			fmt.Println("Ingrese una fruta")
			fmt.Scan(&fruta)
			/*se utiliza el modulo strings para acceder a una serie de metodos, los cuales nos sirven para modificar
			variables de tipo strings: ToLower, ToUpper, Trim entre otros
			*/
			fruta = strings.ToLower(fruta)
		}

	}
}
