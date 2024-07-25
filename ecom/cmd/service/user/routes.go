// Archivo para las rutas de la api
package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// Registro de rutas, se coloca el endpoint y el metodo
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// Funcion para login
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

// Funcion para registro de usuarios
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}
