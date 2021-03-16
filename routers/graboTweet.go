package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/superterro666/twittor/bd"
	"github.com/superterro666/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se a podido insertar el registro"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
