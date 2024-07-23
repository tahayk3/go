package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/PCD/ecom/cmd/service/user"
	"github.com/gorilla/mux"
)

type APISserver struct {
	addr string
	db   *sql.DB
}

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
