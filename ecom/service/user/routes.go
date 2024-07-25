package user

import (
	"net/http"

	"fmt"

	"github.com/PCD/ecom/service/auth"
	"github.com/PCD/ecom/types"
	"github.com/PCD/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
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
	// Obtener el payload JSON
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Validar el payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Carga util invalida", errors))
		return
	}

	// Verificar si el usuario ya existe, en este caso, envia una respuesta y termina con return
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// Hashear la contraseña
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Crear el nuevo usuario
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	// detecta errores
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//Responde con un estado de creado si todo va bien
	utils.WriteJSON(w, http.StatusCreated, nil)

}
