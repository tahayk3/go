package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// Metodo para conectarse a la db
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	//sql.Open abre una conexion a la db
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// Finaliza el programa si hay un error al abrir la conexión
		log.Fatal(err)
	}
	// Devuelve la conexión a la base de datos y nil para el error
	return db, nil
}
