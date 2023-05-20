package routers

import (
	"github.com/gorilla/mux"
	"github.com/giovannirizzolo/webserver/handlers"
)

func InitRouter() *mux.Router{
	rtr := mux.NewRouter()

	rtr.HandleFunc("/animals", handlers.AnimalsHandler)
	rtr.HandleFunc("/animals/{id:[a-zA-Z0-9]+}", handlers.AnimalHandler).Methods("GET")

	return rtr
}