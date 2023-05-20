package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/giovannirizzolo/webserver/routers"
	"github.com/giovannirizzolo/webserver/models"
	"github.com/giovannirizzolo/webserver/middlewares"
)

var animals = []models.Animal{
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
	router := routers.InitRouter()
	router.Use(middlewares.LoggingMiddleware)


	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func createAnimalHandler(w http.ResponseWriter, r *http.Request) {

}




