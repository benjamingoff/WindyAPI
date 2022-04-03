package main

import (
	"github.com/benjamingoff/WeatherAPI/WindyAPI/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// Creation of the router, the routes to the router, and the http listener.
	router := mux.NewRouter()

	routes.WeatherRoutes(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
