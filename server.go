package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/giovannirizzolo/webserver/middlewares"
	"github.com/giovannirizzolo/webserver/models"
	"github.com/giovannirizzolo/webserver/routers"
)

var animals = []models.Animal{
	{
		Name:   "Gionni",
		Breed:  "Canazzo di bancata",
		LegNum: 6,
	},
	{
		Name:   "Jenny",
		Breed:  "Canazzo di bancata",
		LegNum: 4,
	},
}

func main() {
	

	// Define your route handlers before starting the server
	router := routers.InitRouter()
	amw := middlewares.AuthenticationMiddleware{
		TokenUsers: make(map[string]string),
	}

	amw.Populate()

	router.Use(middlewares.LoggingMiddleware)
	router.Use(amw.AuthMiddleware)

	var port = 8080
	var alterantivePort = 3030

	if err := startServer(router, port); err != nil {
		log.Printf("Port %v already in use, fallback to %v", port, alterantivePort)

		if err = startServer(router, alterantivePort); err != nil {
			log.Fatalf("Both ports %v and %v in use , quitting...", port, alterantivePort)
		}
	}
}

func startServer(router http.Handler, port int) error {
	addr := ":" + strconv.Itoa(port)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	fmt.Printf("Starting server at port %v...\n", port)
	if err := server.ListenAndServe(); err != nil {
		return err
	} 

	return nil

}
