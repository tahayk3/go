package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Metodo para hacer hash a  contraseña por medio de bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// funcion para comparar contraseñas(hasheadas en la db)
func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}
