/*
Una goroutine es una función que se ejecuta de manera concurrente con otras goroutines en el mismo espacio de direcciones.
Son más ligeras que los hilos (threads) y están diseñadas para ser fáciles de usar.
*/
package main

import (
	"fmt"
	"time"
)

func imprimirMensaje(mensaje string) {
	for i := 0; i < 5; i++ {
		fmt.Println(mensaje, i)
		/* Al añadir un pequeño retraso en cada iteración del bucle, se simula que la goroutine está realizando algún tipo de
		trabajo que toma tiempo. En un programa real, podrías estar leyendo archivos, consultando una base de datos, o realizando alguna operación
		 compleja que no sea instantánea.*/
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	//Lanza goroutine uno y dos
	go imprimirMensaje("Goroutine uno")
	go imprimirMensaje("Goroutine dos")
	//Permite que las goroutines se ejecuten durante un segundo antes de que el programa termine
	time.Sleep(1 * time.Second)
	fmt.Println("Fin del programa")
}
