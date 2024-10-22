package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context) *mongo.Client {
	// Creates options that include the logger specification
	clientOptions := options.
		Client().
		ApplyURI(Config.DB.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Panicf("An error occurred while connectiong mongodb: %s", err)
	}

	return client
}
