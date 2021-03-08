package routers

import (
	"encoding/json"
	"net/http"

	"github.com/superterro666/twittor/bd"
	"github.com/superterro666/twittor/models"
)

// Registro :  usuarios
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe tener 6 o mas caracteres ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "El email ya esta en uso", 400)
		return
	}

	_, status, err := bd.InsertRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error inesperado"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo insertar", 400)
	}

	w.WriteHeader(http.StatusCreated)

}
