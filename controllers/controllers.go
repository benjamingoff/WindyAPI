package controllers

import (
	"context"
	"encoding/json"
	"github.com/benjamingoff/WeatherAPI/WindyAPI/configs"
	"github.com/benjamingoff/WeatherAPI/WindyAPI/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var weatherCollection *mongo.Collection = configs.GetCollection(configs.DB, "Weather")

func GetMostRecentWeather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

		var weather responses.InnerWeatherObject

		err := weatherCollection.FindOne(ctx, bson.M{"$natural": -1}).Decode(&weather)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(weather)
	}
}
