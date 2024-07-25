package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PCD/ecom/types"
	"github.com/gorilla/mux"
)

// Pruebas para  manejadores del servicio de usuarios
func TestUserServiceHandlers(t *testing.T) {

	//Creando un objeto simulado que imita el comportamiento de objetos reales en un entorno controlado
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	//Se nombra a la prueba
	t.Run("should handle get user by ID", func(t *testing.T) {

		//Se crea una payload para la prueba
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "tahayk3@gmail.com",
			Password:  "asd345",
		}

		//Se convierte el payload en un json
		marshalled, _ := json.Marshal(payload)

		//Creacion de la solicitud http a la ruta /register
		req, err := http.NewRequest(http.
			MethodPost, "/register", bytes.NewBuffer(marshalled))

		//Se interrumpe el proceso en caso de algun error
		if err != nil {
			t.Fatal(err)
		}

		//httptest.NewRecorder() grabar las respuestas HTTP
		rr := httptest.NewRecorder()
		//Creacion de nueva router
		router := mux.NewRouter()
		//Se asocia la ruta /register con el manejador handleRegister.
		router.HandleFunc("/register", handler.handleRegister)
		//Se envía la solicitud al router.
		router.ServeHTTP(rr, req)

		//Se compueba el estado http estperado
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}
