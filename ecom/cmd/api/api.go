package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/PCD/ecom/service/cart"
	"github.com/PCD/ecom/service/order"
	"github.com/PCD/ecom/service/product"
	"github.com/PCD/ecom/service/user"
	"github.com/gorilla/mux"
)

// Estructura para el servidor, addr: puerto en el que escucha el servidor, db: url a la que debe conectarse el servidor
type APIServer struct {
	addr string
	db   *sql.DB
}

// Funcion para crear servidor
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	// Serve para archivos estaticos
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
