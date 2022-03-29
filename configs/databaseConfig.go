package configs

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"time"
)

func ConnectToDatabase() *firestore.Client {

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	sa := option.WithCredentialsFile("C:\\Users\\benja\\go\\src\\github.com\\benjamingoff\\WeatherAPI\\WindyAPI\\configs\\windyapidatabase-84bdca0a7349.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	return client
}

var DB *firestore.Client = ConnectToDatabase()
