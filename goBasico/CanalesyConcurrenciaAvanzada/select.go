package main

import (
	"fmt"
	"time"
)

/*
El select en Go es similar a un switch pero para canales. Te permite esperar en múltiples
operaciones de canal y proceder con la primera que esté lista.
*/
func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		canal1 <- "canal 1 listo"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		canal2 <- "canal 2 listo"
	}()
	for i := 0; i < 2; i++ {
		select {
		case mensaje1 := <-canal1:
			fmt.Println(mensaje1)
		case mensaje2 := <-canal2:
			fmt.Println(mensaje2)
		}
	}
}
