package main

import (
	"log"

	"github.com/superterro666/twittor/bd"
	"github.com/superterro666/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()

}
