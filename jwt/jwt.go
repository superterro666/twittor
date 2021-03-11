package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/superterro666/twittor/models"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MastersDelDesarrolloRolo$$")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,

		"sitioweb": t.SitioWeb,
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
