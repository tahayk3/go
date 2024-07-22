package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
Estructura que emula una tabla en una db, en esta se realizan
operaciones CRUD
*/
type Article struct {
	Title   string `json:"Title"`
	Descr   string `json: "Descr"`
	Content string `json:"Content"`
}

type Articles []Article

//funciones para manipular y recuperar informacion de la estructura
/*Estas funciones suelen necesitar dos parametros: un escritor y un
lector de las solicitudes
*/
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Titulo test", Descr: "Descripcion Test", Content: "Hola mundo"},
	}

	fmt.Println("Endpoint Hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

// endpoint de la pagina principal
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Acceso al punto final de la página principal")
}

// Manejo de solicitudes
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Punto de entrada donde se llama a la funcion que escucha solicitudes
func main() {
	handleRequest()
}
