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
	// Get the credentials from the private key file provided, then spin up a firebase instance.
	sa := option.WithCredentialsFile("controllers/windyapidatabase-firebase-adminsdk-p6hix-46ad1ad13e.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal(err)
	}

	// Create a firestore client.
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the firestore.")
	return client
}

// Global variable that can be accessed where needed
var client = createClient(context.Background())

func GetMostRecentWeather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := context.Background()

		// Query on the collection, grabs the most recent document by timestamp in the document
		query := client.Collection("WeatherReports").OrderBy("Time", firestore.Desc).Limit(1).Documents(ctx)

		// Technically returns an iterator of one object, so just incrementing to the object it gives back.
		doc, err := query.Next()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err, "Unable to fetch most recent weather report.")
		}

		// This doesn't work. TODO: Make it work someday
		log.Println("Request from ", r.Header.Get("X-Forwarded-For"))

		// Returns the documents' data, maybe add some validation here?
		err = json.NewEncoder(w).Encode(doc.Data())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}
