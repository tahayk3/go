/*Ejercicio 3: Contador Concurrente
Crea un programa donde múltiples goroutines incrementen un contador compartido. Usa un canal para sincronizar las goroutines y asegurarte de
que no haya condiciones de carrera (race conditions). Haz que el programa imprima el valor final del contador.
*/

package main

import (
	"fmt"
)

func main() {
	contador := 0
	canal := make(chan int)

	// Función para incrementar el contador
	incrementar := func() {
		for valor := range canal {
			contador += valor
		}
	}

	// Lanzar la goroutine que maneja el incremento
	go incrementar()

	// Enviar incrementos al canal, la gourutine se queda escuchando, cuando el canal cambia, la funcion incrementar, se ejecuta nuevamente
	//tanto el for como la gorutine, se ejecutan al mismo tiempo(concurrente)
	for i := 0; i <= 10; i++ {
		canal <- 1
	}

	// Cerrar el canal
	close(canal)

	fmt.Println("Valor final del contador:", contador)
}
