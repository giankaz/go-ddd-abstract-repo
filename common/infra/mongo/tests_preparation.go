package mongorepo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoTestsPreparation() (*mongo.Collection, *mongo.Client, *mongo.Database) {
	uri := "mongodb://localhost:27017"

	if uri == "" {
		log.Fatal("You must set your 'MONGO' uri")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Println(err)
	}

	database := client.Database("testing-db")

	collection := database.Collection("mock-collection")

	return collection, client, database

}
