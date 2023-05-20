package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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

	router.HandleFunc("/animals", AnimalsHandler).Methods("GET")
	router.HandleFunc("/animals/{id:[a-zA-Z0-9]+}", AnimalHandler).Methods("GET")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func createAnimalHandler(w http.ResponseWriter, r *http.Request) {

}

func AnimalHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Animal id: %v", vars["id"])
	fmt.Fprintf(w, "Animal id: %v", vars["id"])
}

func AnimalsHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	// fmt.Printf("Animals: %v",["id"])
	// fmt.Fprintf(w, "Animal id: %v", vars["id"])
}



