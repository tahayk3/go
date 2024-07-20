package main

import "fmt"

func main() {
	var edad int
	fmt.Println("Ingresa una edad: ")
	//solicitando datos en terminal
	fmt.Scan(&edad)
	//En los ifs, los () no son necesarios
	if edad >= 18 {
		fmt.Println("Persona mayor de edad")
	} else if edad < 18 {
		fmt.Println("La persona es menor de edad")
	} else {
		fmt.Println("Edad no validad")
	}

	//La estructura del switch es la siguiente
	/*
		switch fruta{
			case "opcion1":
				fmt.Println("Edad no validad")
			case "opcion2":
				fmt.Println("Edad no validad")
			case "opcion3":
				fmt.Println("Edad no validad")
			default:
				fmt.Println("Edad no validad")
		}
	*/
}
