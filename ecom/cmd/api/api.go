package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/PCD/ecom/cmd/service/user"
	"github.com/gorilla/mux"
)

// Estructura para el servidor, addr: puerto en el que escucha el servidor, db: url a la que debe conectarse el servidor
type APISserver struct {
	addr string
	db   *sql.DB
}

// Funcion para crear servidor
func NewAPIServer(addr string, db *sql.DB) *APISserver {
	return &APISserver{
		addr: addr,
		db:   db,
	}
}

func (s *APISserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("Servidor escuchando ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
