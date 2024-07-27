package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// En initConfig se utiliza un singleton para evitar la creacion constante de initConfig
var Envs = initConfig()

// Estructura para la configuracion a la conexion a la db
type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

// funcion que utiliza la estructura para crear la conexion a la db
func initConfig() Config {
	godotenv.Load() // Carga las variables de entorno desde un archivo .env
	//Inicializa una estructura Config con valores de entorno o valores por defecto.
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecom"),

		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
	}
}

// Carga variables de entorno desde un archivo .env usando godotenv.Load
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		// Devuelve el valor de la variable de entorno si está configurada
		return value
	}
	// Devuelve el valor por defecto si la variable de entorno no está configurada
	return fallback
}

// Obtener el valor de una variable de entorno y convertirlo a un entero de 64 bits
func getEnvAsInt(key string, fallback int64) int64 {
	//LookupEnv buscar la variable de entorno especificada por key, si no existe, usa la que esta por defecto
	if value, ok := os.LookupEnv(key); ok {
		//Conversion a entero
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
