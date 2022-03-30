package controllers

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func createClient(ctx context.Context) *firestore.Client {
	sa := option.WithCredentialsFile("controllers/windyapidatabase-firebase-adminsdk-p6hix-46ad1ad13e.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the firestore.")
	return client
}

var client = createClient(context.Background())

func GetMostRecentWeather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := context.Background()

		query := client.Collection("WeatherReports").OrderBy("Time", firestore.Desc).Limit(1).Documents(ctx)

		doc, err := query.Next()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err, "Unable to fetch most recent weather report.")

		}

		log.Println("Request from ", r.Header.Get("X-Forwarded-For"))

		err = json.NewEncoder(w).Encode(doc.Data())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}
