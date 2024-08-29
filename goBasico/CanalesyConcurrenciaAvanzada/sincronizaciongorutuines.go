/*
sync.WaitGroup se usa para esperar a que un conjunto de goroutines termine de ejecutarse.
Esto es útil cuando tienes múltiples tareas concurrentes y necesitas esperar a que todas
terminen antes de continuar.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func trabajar(id int, wg *sync.WaitGroup) {
	//Marca que una goroutine ha terminado
	//El método Done se llama al final de cada goroutine, decrementando el contador.
	defer wg.Done()

	fmt.Printf("Goroutine %d esta trabajando\n", id)
	// Simular un trabajo
	time.Sleep(time.Second * 2)
	fmt.Printf("Goroutine %d ha terminado\n", id)
}
func main() {
	var wg sync.WaitGroup

	for x := 0; x <= 5; x++ {
		// Añadir una goroutine al WaitGroup
		//El método Add incrementa el contador del WaitGroup por cada goroutine lanzada.
		wg.Add(1)
		go trabajar(x, &wg)
	}

	//Esperar a que todas las goroutines terminen
	//Wait bloquea la ejecución hasta que el contador llegue a cero.
	wg.Wait()
	fmt.Println("Todas las goroutines han terminado")
}
