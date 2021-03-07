package middlewares

import (
	"net/http"

	"github.com/superterro666/twittor/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la bd", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
