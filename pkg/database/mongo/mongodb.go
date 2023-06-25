package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	ctx    context.Context
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

var DB *mongo.Client = ConnectDB()


func Connect(connectionString, dbName string) error {
	var err error

	// Create a context with a timeout (optional)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configure the MongoDB client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to the MongoDB server
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	return nil
}

func GetDB(ctx context.Context, dbName string) *mongo.Database {
	return client.Database(dbName)
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("moviedb").Collection(collectionName)
	return collection
}

func InsertDocument(dbName string, collectionName string, document interface{}) error {
	collection := client.Database(dbName).Collection(collectionName)

	// Insert the document
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}

	return nil
}
