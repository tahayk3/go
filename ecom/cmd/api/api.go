package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/PCD/ecom/service/product"
	"github.com/PCD/ecom/service/user"
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

	//usuarios
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	//productos
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
