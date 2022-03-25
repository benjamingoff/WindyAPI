package controllers

import (
	"context"
	"encoding/json"
	"github.com/benjamingoff/WeatherAPI/WindyAPI/configs"
	"github.com/benjamingoff/WeatherAPI/WindyAPI/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

var weatherCollection *mongo.Collection = configs.GetCollection(configs.DB, "Weather")

func GetMostRecentWeather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := context.Background()

		var weather responses.InnerWeatherObject

		myOptions := options.FindOne()
		myOptions.SetSort(bson.M{"$natural": -1})

		err := weatherCollection.FindOne(ctx, bson.M{}, myOptions).Decode(&weather)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(weather)
	}
}
