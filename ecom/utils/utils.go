package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// Decodifica el cuerpo de la solicitud JSON en el objeto payload proporcionado.
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

// Codifica el objeto v en JSON y lo escribe en la respuesta con el estado HTTP dado.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// Funcion para mostrar errores
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
