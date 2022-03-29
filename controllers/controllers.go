package controllers

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/benjamingoff/WeatherAPI/WindyAPI/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

var client = configs.DB

func GetMostRecentWeather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := context.Background()

		// var weather responses.InnerWeatherObject

		myOptions := options.FindOne()
		myOptions.SetSort(bson.M{"$natural": -1})

		query := client.Collection("WeatherReports").OrderBy("Time", firestore.Desc).Limit(1).Documents(ctx)

		doc, _ := query.Next()
		json.NewEncoder(w).Encode(doc.Data())
		w.WriteHeader(http.StatusOK)

	}
}
