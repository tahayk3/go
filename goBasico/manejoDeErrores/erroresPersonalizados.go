package main

import (
	"fmt"
)

type ErrorPersonalizado struct {
	mensaje string
}

func (e *ErrorPersonalizado) Error() string {
	return e.mensaje
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &ErrorPersonalizado{mensaje: "División por cero detectada"}
	}
	return a / b, nil
}

func main() {
	_, err := dividir(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
