package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"example.com/handlers"
)

type Animal struct {
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	LegNum int    `json:"legNum"`
}


var animals = []Animal{
	{
		Name: "Gionni",
		Breed: "Canazzo di bancata",
		LegNum: 6,
	},
	{
		Name: "Jenny",
		Breed: "Canazzo di bancata",
		LegNum: 4,
	},
}

func main() {
	fmt.Println("Starting server at port 8080")

	// Define your route handlers before starting the server
	router := mux.NewRouter()

	router.HandleFunc("/animals", handlers.AnimalsHandler).Methods("GET")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func createAnimalHandler(w http.ResponseWriter, r *http.Request) {

}




