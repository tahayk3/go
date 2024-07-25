package main

import (
	"database/sql"
	"log"

	"github.com/PCD/ecom/cmd/api"
	"github.com/PCD/ecom/config"
	"github.com/PCD/ecom/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// Configura la conexión a la base de datos usando valores del archivo .env o valores por defecto
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Inicializa  base de datos, comprueba la conexion
	initStorage(db)

	// Crea un nuevo servidor API en el puerto 8080
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// initStorage verifica la conexión a la base de datos
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Conexion a la base de datos exitosa!")
}
