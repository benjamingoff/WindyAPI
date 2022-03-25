package routes

import (
	"github.com/benjamingoff/WeatherAPI/WindyAPI/controllers"
	"github.com/gorilla/mux"
)

func WeatherRoutes(router *mux.Router) {
	router.HandleFunc("/api/mostRecentWeather", controllers.GetMostRecentWeather()).Methods("GET")
}
