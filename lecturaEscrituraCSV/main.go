package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Estructura para guardar datos de una archivo
type Alumno struct {
	Nombre string
	Notas  []float64
}

// Funcion para la lectura de una archivo CSV
func readCSV(nombreArchivo string) ([]Alumno, error) {
	//Lectura de archivo  y se realiza la validacion en caso de error
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}

	//Agregar ; al momento de leer el archivo
	lector := csv.NewReader(archivo)
	lector.Comma = ';'

	//Se guardan los registros y se realiza la validacion en caso de error
	registros, err := lector.ReadAll()
	if err != nil {
		return nil, err
	}

	var alumnos []Alumno

	//Utilizando ciclos para leer todo el archivo y para ingresar los datos leidos en la estructura alumnos
	for _, registro := range registros {
		nombre := registro[0]
		var notas []float64
		for _, notasStr := range registro[1:] {
			//Conversion de notas, de string a float
			nota, err := strconv.ParseFloat(notasStr, 64)
			if err != nil {
				return nil, err
			}
			notas = append(notas, nota)
		}
		alumnos = append(alumnos, Alumno{Nombre: nombre, Notas: notas})
	}
	return alumnos, nil
}

func calcularPromedio(notas []float64) float64 {
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

func calcularMediaAritmentica(alumnos []Alumno) float64 {
	var suma float64
	var cantidadNotas int
	for _, alumno := range alumnos {
		for _, nota := range alumno.Notas {
			suma = suma + nota
			cantidadNotas++
		}
	}
	return suma / float64(cantidadNotas)
}

// funcion para escribir datos en un archivo
func writeCSV(nombreArchivo string, alumnos []Alumno) error {
	//crear archivo y manejo de error
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return err
	}

	//Cierre de archivo
	defer archivo.Close()

	//estructura de tipo csv para el nuevo archivo
	escritor := csv.NewWriter(archivo)
	//liberar memoria
	defer escritor.Flush()

	//escritura para cada alumno en el archivo
	for _, alumno := range alumnos {
		//En base a la funcion calcularPromedio, obtenemos el promedio
		promedio := calcularPromedio(alumno.Notas)
		//Escribimos el nombre del alumno, junto con su promedio
		registro := []string{alumno.Nombre, fmt.Sprintf("%.2f", promedio)}
		//Manejo de errores
		if err := escritor.Write(registro); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	alumnos, err := readCSV("notas.csv")

	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	//Media aritmentica
	mediaAritmentica := calcularMediaAritmentica(alumnos)
	fmt.Printf("La media aritmentica de todas las notas es: %.2f\n", mediaAritmentica)

	//escritura de nuevo archivo
	err = writeCSV("promedios.csv", alumnos)
	if err != nil {
		fmt.Println("Error al escribir el archivo CSV", err)
		return
	}

	fmt.Println("Archivo creado con exitosamente")

}
