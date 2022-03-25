package main

import (
	"github.com/benjamingoff/WeatherAPI/WindyAPI/configs"
	"github.com/benjamingoff/WeatherAPI/WindyAPI/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	routes.WeatherRoutes(router)

	configs.ConnectToDatabase()

	log.Fatal(http.ListenAndServe(":8000", router))
}
