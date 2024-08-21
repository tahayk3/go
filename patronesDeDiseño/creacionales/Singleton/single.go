package main

import (
	"fmt"
	"sync"
)

/*
		lock es un puntero a un Mutex de la librería sync, utilizado para asegurar que solo un goroutine pueda ejecutar una sección
	 	crítica del código a la vez. Esdto evita la condicion de carrera
*/
var lock = &sync.Mutex{}

// estructura vacía que representa el tipo de la instancia única que queremos crear.
type single struct{}

// singleInstance es un puntero a la estructura single, y es donde se almacenará la única instancia de esta estructura.
var singleInstance *single

func getInstance() *single {
	//Verifica que la instancia no ha sido creada
	if singleInstance == nil {
		//de no estar creada, se bloquea a todas las demas gorutines para que solo una pueda tener acceso a esa seccion critica de codigo y pueda crear el singleton
		lock.Lock()
		//defer lock.Unlock() asegura que el bloqueo se libere cuando la función termine.
		defer lock.Unlock()
		//se vuelve a verificar si el singleton(instancia) no ha sido creada por otra gouroutine durante el proceso anterior
		if singleInstance == nil {
			fmt.Println("creando la instancia ahora")
			//se crea la instancia
			singleInstance = &single{}
		} else {
			fmt.Println("La instancia ya fue creada")
		}
	} else {
		fmt.Println("La instancia ya fue creada")
	}
	return singleInstance
}
