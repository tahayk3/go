/*Sin excepciones: A diferencia de muchos otros lenguajes, Go no utiliza excepciones para manejar errores. En su lugar, las funciones suelen
 devolver dos valores: el resultado esperado y un valor de tipo error.

Error como valor: En Go, un error es un valor como cualquier otro, lo que significa que puede ser asignado a variables, pasado como argumento
 a funciones y comparado.


Comprobación explícita: El programador debe comprobar explícitamente si se ha producido un error y tomar las medidas adecuadas. Esto puede
parecer más verboso al principio, pero fomenta un estilo de programación más seguro y predecible.

Muchas funciones en Go devuelven dos valores: el resultado esperado y un valor de tipo error.
El valor de error es nil si no se ha producido ningún error.
*/

package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", result)
	}
}
