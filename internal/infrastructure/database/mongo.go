package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Initialize the database connection
func Connect(ctx context.Context) (err error) {

	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return
	}

	log.Println("Connected to MongoDB!")

	return

}

// GetClient returns the MongoDB client.
func GetClient() *mongo.Client {
	return client
}

// Get the database
func GetDatabase() *mongo.Database {
	return client.Database(os.Getenv("MONGO_DATABASE"))
}

// close the database connection
func Disconnect(ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}
