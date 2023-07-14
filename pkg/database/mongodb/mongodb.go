package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Mhmdaris15/booking-movie-app/internal/configs"
	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	ctx    context.Context
)

func ConnectDB() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(configs.EnvMongoURI()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
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

func DisconnectDB(client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
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

func SeedingDatabase(client *mongo.Client) ([]models.User, []models.Cinema, []models.Showtime, []models.Seat) {
	users := models.SeedUser(GetCollection(client, "users"))
	cinemas := models.SeedCinema(GetCollection(client, "cinema"))
	showtimes := models.SeedShowtime(GetCollection(client, "showtime"), GetCollection(client, "movie"), GetCollection(client, "cinema"), GetCollection(client, "seat"))
	// seats := models.SeedSeat(GetCollection(client, "seat"), GetCollection(client, "showtime"))
	seats := []models.Seat{}
	return users, cinemas, showtimes, seats
}
