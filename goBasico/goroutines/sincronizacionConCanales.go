package main

import "fmt"

/*
Los canales (channels) son una forma de comunicación entre goroutines.
Permiten pasar datos de una goroutine a otra de manera segura.
*/

//resultado chan int, indica que va a trabajar con un canal
func sumar(a int, b int, resultado chan int) {
	//Enviar resultado al canal
	resultado <- a + b
}

func main() {
	//se crea un canal, se indica que tipo de variable va a trabajar el canal(envia y recibe enteros)
	resultado := make(chan int)

	//se crea una goroutine que llama a la funcion sumar
	go sumar(3, 4, resultado)

	//Extraer el resultado desde el canal
	suma := <-resultado
	fmt.Println("El resultado extraido del canal es: ", suma)
}
