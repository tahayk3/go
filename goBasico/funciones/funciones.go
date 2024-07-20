package main

import "fmt"

func sumar(num1 int, num2 int) int {
	var suma int
	suma = num1 + num2
	return suma
}
func main() {
	var resultado int = sumar(5, 5)
	fmt.Println(resultado)
}
