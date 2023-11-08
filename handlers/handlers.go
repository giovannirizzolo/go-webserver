package handlers

import (
	
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)	

	// data, err := w.Write(json)

	// fmt.Print(data)
	// fmt.Print(err)
}