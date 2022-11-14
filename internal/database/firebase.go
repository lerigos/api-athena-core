package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
)

func QueryPingMonitors(client *firestore.Client) {
	opt := option.WithCredentialsFile("../saygames.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	collection := client.Collection("monitors")
	query := collection.OrderBy("monitor_owner", firestore.Desc).Where("monitor_type", "==", "ping")

	_ = query
	print(query)
}

func QueryWebMonitors(client *firestore.Client) {
	opt := option.WithCredentialsFile("../saygames.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	collection := client.Collection("monitors")
	query := collection.OrderBy("monitor_owner", firestore.Desc).Where("monitor_type", "==", "http")

	_ = query
	print(query)
}
