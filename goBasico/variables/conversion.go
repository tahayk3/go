package main

import "fmt"

var a int

type dinero int

var b dinero

func main() {
	a = 42
	fmt.Println(a)
	//mostrar tipo
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	//Conversion de b a int
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
