package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/superterro666/twittor/middlewares"
	"github.com/superterro666/twittor/routers"
)

/* Manejador seteo puerto y handler y escuchamos el servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.GraboTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
