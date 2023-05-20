package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"	
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