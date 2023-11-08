package routers

import (
	"net/http"

	"github.com/giovannirizzolo/webserver/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router{
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", handlers.HomeHandler)
	rtr.HandleFunc("/animals", handlers.AnimalsHandler)
	rtr.HandleFunc("/animals/{id:[a-zA-Z0-9]+}", handlers.AnimalHandler)
	http.Handle("/", rtr)

	return rtr
}