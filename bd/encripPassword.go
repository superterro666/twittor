package bd

import (
	"golang.org/x/crypto/bcrypt"
)

// EncriptarPassword : encriptar
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	byte, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(byte), err
}
