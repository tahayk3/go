package auth

import (
	"strconv"
	"time"

	"github.com/PCD/ecom/config"
	"github.com/golang-jwt/jwt/v5"
)

/*
Parametros
Secret: es un arreglo de bytes que representa la clave secreta utilizada para firmar el token.
userID: es el ID del usuario para quien se está creando el token
*/
func CreateJWT(secret []byte, userID int) (string, error) {
	//duracion del token en segundos por medio de  JWTExpirationInSeconds
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	//Creacion del token JWT por medio de los parametros enviados y el metodo HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//Aclaraciones relacionados con el usuario: id del usuario y tiempo de vida el token
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	//Se firma el token con la clave secreta proporcionada
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
