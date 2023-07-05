package models

import (
	"context"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Showtime struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MovieID   primitive.ObjectID `bson:"movie_id,omitempty" json:"movie_id,omitempty"`
	CinemaID  primitive.ObjectID `bson:"cinema_id,omitempty" json:"cinema_id,omitempty"`
	StartTime string
	EndTime   string
}

func SeedShowtime(showtimeCollection *mongo.Collection, movieCollection *mongo.Collection ,cinemaCollection *mongo.Collection) ([]Showtime) {
	var showtimes []Showtime
	showtimeCount := 5

	// Retrieve all ID from movie collection
	var movies []Movie
	cursor, err  := movieCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.Background(), &movies); err != nil {
		log.Fatal(err)
	}

	// Retrieve all ID from cinema collection
	var cinemas []Cinema
	cursor, err  = cinemaCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.Background(), &cinemas); err != nil {
		log.Fatal(err)
	}

	movieIDs := []primitive.ObjectID{}
	cinemaIDs := []primitive.ObjectID{}

	for _, movie := range movies {
		movieIDs = append(movieIDs, movie.ID)
	}

	for _, cinema := range cinemas {
		cinemaIDs = append(cinemaIDs, cinema.ID)
	}



	for i := 0; i < showtimeCount; i++ {
		// Get Random Movie ID
		randomMovieID := movieIDs[rand.Intn(len(movieIDs))]
		// Get Random Cinema ID
		randomCinemaID := cinemaIDs[rand.Intn(len(cinemaIDs))]

		showtime := Showtime{
			ID: 	 primitive.NewObjectID(),
			MovieID: randomMovieID,
			CinemaID: randomCinemaID,
			StartTime: faker.TimeString(),
			EndTime: faker.TimeString(),
		}
		showtimes = append(showtimes, showtime)

		_, err := showtimeCollection.InsertOne(context.Background(), showtime)
		if err != nil {
			log.Fatal(err)
		}
	}

	return showtimes
}

