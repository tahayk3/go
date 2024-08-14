package main

import (
	"fmt"
	"time"
)

/*
Un canal con buffer permite almacenar valores antes de que sean recibidos, lo cual puede ser útil para evitar bloqueos si
una goroutine no está lista para recibir datos de inmediato.
*/

func productor(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Enviando:", i)
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumidor(ch chan int) {
	for v := range ch {
		fmt.Println("Recibiendo:", v)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// Creamos un canal con un buffer de tamaño 2, esto evita que la goroutine productor, no se bloquee ya que es mas rapida que el consumidor
	ch := make(chan int, 2)

	go productor(ch)
	go consumidor(ch)

	time.Sleep(time.Second * 11)
}
