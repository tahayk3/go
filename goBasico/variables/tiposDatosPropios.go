package main

import "fmt"

//para crear nuestro propio tipo de dato, este debe basarse en uno primitivo(int en este caso)
type dinero int

var b dinero

func main() {
	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
